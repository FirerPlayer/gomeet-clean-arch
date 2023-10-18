package web

import (
	"github.com/firerplayer/whatsmeet-go/internal/infra/web/webserver"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
	usecase "github.com/firerplayer/whatsmeet-go/internal/usecase/message"
	"github.com/gofiber/fiber/v2"
)

type MessageWebHandlers struct {
	webServer *webserver.WebServer
	*usecase.ListMessageByChatIDUsecase
}

func NewMessageWebHandlers(
	wb *webserver.WebServer,
	listMessageByChatIDUsecase *usecase.ListMessageByChatIDUsecase,
) *MessageWebHandlers {
	return &MessageWebHandlers{
		webServer:                  wb,
		ListMessageByChatIDUsecase: listMessageByChatIDUsecase,
	}
}

func (u *MessageWebHandlers) RegisterRoutes() {
	u.webServer.Get("/message/all/:chatID", u.ListMessageByChatID)
}

func (u *MessageWebHandlers) ListMessageByChatID(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	chatID := c.Params("chatID", "invalid")
	if chatID == "invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "chatID is required",
		})
	}
	messageOut, err := u.ListMessageByChatIDUsecase.Execute(c.Context(), dto.ListMessageByChatIDInputDTO{ChatID: chatID, Limit: limit})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(messageOut)
}
