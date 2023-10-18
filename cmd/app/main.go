package main

import (
	"context"
	"errors"
	"log"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
	"github.com/firerplayer/whatsmeet-go/internal/infra/repository"
	"github.com/firerplayer/whatsmeet-go/internal/infra/web"
	"github.com/firerplayer/whatsmeet-go/internal/infra/web/webserver"
	"github.com/firerplayer/whatsmeet-go/internal/infra/web/websocket"
	chatUc "github.com/firerplayer/whatsmeet-go/internal/usecase/chat"
	messageUc "github.com/firerplayer/whatsmeet-go/internal/usecase/message"
	userUc "github.com/firerplayer/whatsmeet-go/internal/usecase/user"
)

func main() {
	ctx := context.Background()
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		panic(err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "root"),
	})
	if err != nil {
		panic(err)
	}

	err = testArangoDBServer(ctx, client)
	if err != nil {
		panic(err)
	}

	initializer := arangodb.NewDBInitializer(ctx, client)

	db, err := initializer.Init("gomeet")
	if err != nil {
		panic(err)
	}

	wb := webserver.NewWebServer("8080", "Gomeet")

	ur := repository.NewUserRepository(db, "User")
	cr := repository.NewChatRepository(db, "Chat")
	mr := repository.NewMessageRepository(db, "Message")
	userHandlers := getUserHandlers(wb, ur)
	userHandlers.RegisterRoutes()
	chatHandlers := getChatHandlers(wb, cr, mr)
	chatHandlers.RegisterRoutes()
	messageHandlers := getMessageHandlers(wb, mr)
	messageHandlers.RegisterRoutes()

	hub := websocket.NewHub("/chat/ws", wb, ctx)
	hub.Run()

	if err = wb.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}

func getUserHandlers(wb *webserver.WebServer, uG *repository.UserRepository) *web.UsersWebHandlers {

	createUserUsecase := userUc.NewCreateUserUsecase(uG)
	deleteByIDUsecase := userUc.NewDeleteByIDUsecase(uG)
	// usecase.GetByIDUsecase
	getByIDUsecase := userUc.NewGetByIDUsecase(uG)
	// usecase.GetAllLimitUsersUsecase
	getAllLimitUsersUsecase := userUc.NewGetAllLimitUsersUsecase(uG)
	// usecase.GetByEmailUsecase
	getByEmailUsecase := userUc.NewGetByEmailUsecase(uG)
	// usecase.UpdateByIDUsecase
	updateByIDUsecase := userUc.NewUpdateByIDUsecase(uG)

	handlers := web.NewUsersWebHandlers(wb,
		createUserUsecase,
		deleteByIDUsecase,
		getByIDUsecase,
		getAllLimitUsersUsecase,
		getByEmailUsecase,
		updateByIDUsecase,
	)
	return handlers
}

func getChatHandlers(wb *webserver.WebServer, cR *repository.ChatRepository, mR *repository.MessageRepository) *web.ChatWebHandlers {
	// *chatUc.DeleteChatByIDUsecase
	deleteChatByIDUsecase := chatUc.NewDeleteChatByIDUsecase(cR, mR)
	// *chatUc.GetByChatIDUsecase
	getByChatIDUsecase := chatUc.NewGetByChatIDUsecase(cR)
	// *chatUc.ListChatByUserIDUsecase
	listChatByUserIDUsecase := chatUc.NewListChatByUserIDUsecase(cR)
	// *chatUc.AddUserByChatIDUsecase
	addUserByChatIDUsecase := chatUc.NewAddUserByChatIDUsecase(cR)

	handlers := web.NewChatWebHandlers(wb,
		deleteChatByIDUsecase,
		getByChatIDUsecase,
		listChatByUserIDUsecase,
		addUserByChatIDUsecase,
	)
	return handlers
}

func getMessageHandlers(wb *webserver.WebServer, mR *repository.MessageRepository) *web.MessageWebHandlers {
	listMessageByChatIDUsecase := messageUc.NewListMessageByChatIDUsecase(mR)
	handlers := web.NewMessageWebHandlers(wb, listMessageByChatIDUsecase)
	return handlers
}

func testArangoDBServer(ctx context.Context, client driver.Client) error {
	info, err := client.Version(ctx)
	if err != nil {
		return errors.New("Could not connect to ArangoDB, is it running? Error details: " + err.Error())
	}
	log.Println(info.String())
	return nil
}
