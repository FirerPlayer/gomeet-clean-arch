package websocket

import (
	"context"
	"log"
	"sync"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/web/webserver"
	usecase "github.com/firerplayer/whatsmeet-go/internal/usecase/chat"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
	message "github.com/firerplayer/whatsmeet-go/internal/usecase/message"
	"github.com/gofiber/contrib/websocket"
)

// type websocketConnection struct {
// 	Connection *websocket.Conn
// 	Id string
// }

type chatWS struct {
	ChatID     string
	Connection *websocket.Conn
	mu         sync.Mutex
	isClosing  bool
}

func newChatWS(chat_id string, connection *websocket.Conn) *chatWS {
	return &chatWS{
		ChatID:     chat_id,
		Connection: connection,
	}
}

type Hub struct {
	webServer            *webserver.WebServer
	ctx                  context.Context
	path                 string
	chats                map[*websocket.Conn]*chatWS
	register             chan *websocket.Conn
	broadcast            chan entity.Message
	unregister           chan *websocket.Conn
	createChatUsecase    usecase.CreateChatUsecase
	createMessageUsecase message.CreateMessageUsecase
}

func NewHub(path string, webServer *webserver.WebServer, ctx context.Context) *Hub {
	return &Hub{
		webServer:  webServer,
		ctx:        ctx,
		path:       path,
		chats:      make(map[*websocket.Conn]*chatWS),
		register:   make(chan *websocket.Conn),
		broadcast:  make(chan entity.Message),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) runHub() {
	for {
		select {
		case connection := <-h.register:
			var input dto.CreateChatInputDTO
			err := connection.ReadJSON(&input)
			if err != nil {
				log.Println("read error at conn initialization: ", err)
			}
			out, err := h.createChatUsecase.Execute(h.ctx, input)
			if err != nil {
				log.Println("read error at create chat: ", err)
			}
			wsChat := newChatWS(out.ChatID, connection)
			h.chats[connection] = wsChat
			log.Printf("Websocket connection registered. Chat ID: %s", wsChat.ChatID)

		case message := <-h.broadcast:
			log.Printf("message received: %s", message)

			// Send the message to each connected client
			for connection, chat := range h.chats {
				go func(connection *websocket.Conn, currChat *chatWS) {
					// Acquire a lock on the ChatWS to prevent concurrent modification
					currChat.mu.Lock()
					defer currChat.mu.Unlock()

					// If the ChatWS is being closed, skip sending the message
					if currChat.isClosing {
						return
					}

					// If the ChatWS matches the message's ChatID, send the message
					if currChat.ChatID == message.ChatId {
						newMessage, err := h.createMessageUsecase.Execute(h.ctx, dto.CreateMessageInputDTO{
							ChatID:  currChat.ChatID,
							Content: message.Content,
							File:    message.File,
						})
						if err != nil {
							log.Printf("websocket from chat %s write error: %v", currChat.ChatID, err)
							return
						}

						err = connection.WriteJSON(newMessage)
						// err = connection.WriteMessage(websocket.TextMessage, []byte(message.Content))
						if err != nil {
							// Mark the ChatWS as closing and log the error
							currChat.isClosing = true
							log.Printf("websocket from chat %s write error: %v", currChat.ChatID, err)

							// Send a close message to the client and close the connection
							connection.WriteMessage(websocket.CloseMessage, []byte{})
							connection.Close()

							// Unregister the connection from the hub
							h.unregister <- connection
						}
					}
				}(connection, chat)
			}

		case connection := <-h.unregister:
			// Remove the client from the hub
			delete(h.chats, connection)

			log.Println("connection unregistered")
		}
	}
}

func (h *Hub) Run() {
	go h.runHub()
	log.Println("Running hub and websocket at: /api" + h.path)

	h.webServer.Post(h.path, websocket.New(func(cw *websocket.Conn) {
		// When the function returns, unregister the client and close the connection
		defer func() {
			h.unregister <- cw
			cw.Close()
		}()

		// Register the client
		h.register <- cw

		for {
			var msg entity.Message
			err := cw.ReadJSON(msg)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("read error: %v", err)
				}
				log.Printf("json unmarshal error: %v", err)
				return
			}

			h.broadcast <- msg
		}
	}))
}
