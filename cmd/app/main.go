package main

import (
	"log"

	"github.com/firerplayer/whatsmeet-go/internal/infra/web/webserver"
	"github.com/gofiber/fiber/v2"
)

func main() {

	server := webserver.NewWebServer("8080")
	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})
	log.Fatal(server.Start())

}
