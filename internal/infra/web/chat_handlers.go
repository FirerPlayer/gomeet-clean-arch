package web

import (
	"github.com/firerplayer/whatsmeet-go/internal/infra/web/webserver"
	usecase "github.com/firerplayer/whatsmeet-go/internal/usecase/chat"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
	"github.com/gofiber/fiber/v2"
)

type ChatWebHandlers struct {
	webServer *webserver.WebServer
	*usecase.DeleteChatByIDUsecase
	*usecase.GetByChatIDUsecase
	*usecase.ListChatByUserIDUsecase
	*usecase.AddUserByChatIDUsecase
}

func NewChatWebHandlers(
	wb *webserver.WebServer,

	deleteChatByIDUsecase *usecase.DeleteChatByIDUsecase,
	getByChatIDUsecase *usecase.GetByChatIDUsecase,
	listChatByUserIDUsecase *usecase.ListChatByUserIDUsecase,
	addUserByChatIDUsecase *usecase.AddUserByChatIDUsecase,
) *ChatWebHandlers {
	return &ChatWebHandlers{
		webServer:               wb,
		DeleteChatByIDUsecase:   deleteChatByIDUsecase,
		GetByChatIDUsecase:      getByChatIDUsecase,
		ListChatByUserIDUsecase: listChatByUserIDUsecase,
		AddUserByChatIDUsecase:  addUserByChatIDUsecase,
	}
}

func (u *ChatWebHandlers) RegisterRoutes() {
	u.webServer.Get("/chat/details", u.GetByChatID)
	u.webServer.Get("/chat/all/:userID", u.ListChatByUserID)
	u.webServer.Post("/chat/add-user", u.AddUserByChatID)
	u.webServer.Delete("/chat", u.DeleteByID)
}

// GetByChatID handles the request to retrieve a chat by its ID.
//
// It takes a fiber.Ctx object as a parameter and returns an error.
func (u *ChatWebHandlers) GetByChatID(c *fiber.Ctx) error {
	chatID := c.Query("id")
	if chatID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	chatOut, err := u.GetByChatIDUsecase.Execute(c.Context(), dto.GetByChatIDInputDTO{ChatID: chatID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(chatOut)
}

// ListChatByUserID handles the request to list chats by user ID.
//
// It takes a fiber.Ctx object as a parameter and returns an error.
func (u *ChatWebHandlers) ListChatByUserID(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	userID := c.Params("userID", "invalid")
	if userID == "invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "userID is required",
		})
	}
	chatsOut, err := u.ListChatByUserIDUsecase.Execute(c.Context(), dto.ListChatByUserIDInputDTO{UserID: userID, Limit: limit})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(chatsOut)
}

// AddUserByChatID handles the addition of a user by chatID.
//
// It takes a *fiber.Ctx parameter and returns an error.
func (u *ChatWebHandlers) AddUserByChatID(c *fiber.Ctx) error {
	var input dto.AddUserByChatIDInputDTO
	err := c.BodyParser(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	chat, err := u.AddUserByChatIDUsecase.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(chat)
}

func (u *ChatWebHandlers) DeleteByID(c *fiber.Ctx) error {
	chatID := c.Query("id")
	if chatID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "chatId is required",
		})
	}
	err := u.DeleteChatByIDUsecase.Execute(c.Context(), dto.DeleteChatByIDInputDTO{
		ChatID: chatID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
