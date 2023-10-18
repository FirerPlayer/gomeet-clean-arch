package web

import (
	"github.com/firerplayer/whatsmeet-go/internal/infra/web/webserver"
	"github.com/firerplayer/whatsmeet-go/internal/usecase/dto"
	usecase "github.com/firerplayer/whatsmeet-go/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type UsersWebHandlers struct {
	webServer *webserver.WebServer
	*usecase.CreateUserUsecase
	*usecase.DeleteByIDUsecase
	*usecase.GetByIDUsecase
	*usecase.GetAllLimitUsersUsecase
	*usecase.GetByEmailUsecase
	*usecase.UpdateByIDUsecase
}

func NewUsersWebHandlers(
	wb *webserver.WebServer,
	createUserUsecase *usecase.CreateUserUsecase,
	deleteByIDUsecase *usecase.DeleteByIDUsecase,
	getByIDUsecase *usecase.GetByIDUsecase,
	getAllLimitUsersUsecase *usecase.GetAllLimitUsersUsecase,
	getByEmailUsecase *usecase.GetByEmailUsecase,
	updateByIDUsecase *usecase.UpdateByIDUsecase,
) *UsersWebHandlers {
	return &UsersWebHandlers{
		webServer:               wb,
		CreateUserUsecase:       createUserUsecase,
		DeleteByIDUsecase:       deleteByIDUsecase,
		GetByIDUsecase:          getByIDUsecase,
		GetAllLimitUsersUsecase: getAllLimitUsersUsecase,
		GetByEmailUsecase:       getByEmailUsecase,
		UpdateByIDUsecase:       updateByIDUsecase,
	}
}

func (u *UsersWebHandlers) RegisterRoutes() {
	u.webServer.Get("/user/all", u.GetAllLimitUsers)
	u.webServer.Get("/user/:id", u.GetByID)
	u.webServer.Post("/user", u.CreateUser)
	u.webServer.Delete("/user/:id", u.DeleteByID)
	u.webServer.Put("/user/:id", u.UpdateByID)
}

func (u *UsersWebHandlers) GetByID(c *fiber.Ctx) error {
	id := c.Params("id", "invalid")
	if id == "invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	usrOut, err := u.GetByIDUsecase.Execute(c.Context(), dto.GetUserByIDInputDTO{ID: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(usrOut)

}

func (u *UsersWebHandlers) GetAllLimitUsers(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	usrOut, err := u.GetAllLimitUsersUsecase.Execute(c.Context(), dto.GetAllLimitUsersInputDTO{
		Limit: limit,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(usrOut)
}

func (u *UsersWebHandlers) CreateUser(c *fiber.Ctx) error {
	var input dto.CreateUserInputDTO
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	usrOut, err := u.CreateUserUsecase.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(usrOut)
}

func (u *UsersWebHandlers) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id", "invalid")
	if id == "invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	err := u.DeleteByIDUsecase.Execute(c.Context(), dto.DeleteUserByIDInputDTO{
		ID: id,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (u *UsersWebHandlers) UpdateByID(c *fiber.Ctx) error {
	id := c.Params("id", "invalid")
	if id == "invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
		})
	}
	var input dto.UpdateUserByIDInputDTO
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	out, err := u.UpdateByIDUsecase.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(out)

}
