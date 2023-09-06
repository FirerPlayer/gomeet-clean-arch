package websocket

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// type websocketConnection struct {
// 	Connection *websocket.Conn
// 	Id string
// }

type ChatWS struct {
	entity.Chat
	Connection *websocket.Conn
	mu         sync.Mutex
	isClosing  bool
}

func NewChatWS(chat entity.Chat, connection *websocket.Conn) *ChatWS {
	return &ChatWS{
		Chat:       chat,
		Connection: connection,
	}
}

type Hub struct {
	app        *fiber.App
	ctx        *fiber.Ctx
	path       string
	chats      map[*websocket.Conn]*ChatWS
	register   chan *websocket.Conn
	broadcast  chan entity.Message
	unregister chan *websocket.Conn
}

func NewHub(path string, app *fiber.App, ctx *fiber.Ctx) *Hub {
	return &Hub{
		app:        app,
		ctx:        ctx,
		path:       path,
		chats:      make(map[*websocket.Conn]*ChatWS),
		register:   make(chan *websocket.Conn),
		broadcast:  make(chan entity.Message),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) runHub() {
	for {
		select {
		case connection := <-h.register:
			var chat entity.Chat

			// Read the incoming JSON message from the connection
			err := connection.ReadJSON(chat)
			if err != nil {
				log.Println("read error at conn initialization: ", err)
			}

			// Create a new ChatWS instance and register the connection
			h.chats[connection] = NewChatWS(chat, connection)
			log.Printf("Websocket connection registered. Chat ID: %s", chat.Id.String())

		case message := <-h.broadcast:
			log.Printf("message received: %s", message)

			// Send the message to each connected client
			for connection, chat := range h.chats {
				go func(connection *websocket.Conn, currChat *ChatWS) {
					// Acquire a lock on the ChatWS to prevent concurrent modification
					currChat.mu.Lock()
					defer currChat.mu.Unlock()

					// If the ChatWS is being closed, skip sending the message
					if currChat.isClosing {
						return
					}

					// If the ChatWS matches the message's ChatID, send the message
					if currChat.Id.String() == message.ChatId.String() {
						err := connection.WriteMessage(websocket.TextMessage, []byte(message.Content))
						if err != nil {
							// Mark the ChatWS as closing and log the error
							currChat.isClosing = true
							log.Printf("websocket from user %s write error: %v", currChat.Id.String(), err)

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

	h.app.Post(h.path, websocket.New(func(cw *websocket.Conn) {
		// When the function returns, unregister the client and close the connection
		defer func() {
			h.unregister <- cw
			cw.Close()
		}()

		// Register the client
		h.register <- cw

		for {
			_, message, err := cw.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("read error: %v", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			var msg entity.Message

			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Printf("json unmarshal error: %v", err)
			}

			h.broadcast <- msg
		}
	}))
}
