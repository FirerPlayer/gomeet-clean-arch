package webserver

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberHandler func(c *fiber.Ctx) error

type WebServer struct {
	server         *fiber.App
	getHandlers    map[string]FiberHandler
	postHandlers   map[string]FiberHandler
	deleteHandlers map[string]FiberHandler
	putHandlers    map[string]FiberHandler
	webServerPort  string
}

func NewWebServer(port, appName string) *WebServer {
	return &WebServer{
		server: fiber.New(fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
			AppName:     appName,
		}),
		getHandlers:    make(map[string]FiberHandler),
		postHandlers:   make(map[string]FiberHandler),
		deleteHandlers: make(map[string]FiberHandler),
		putHandlers:    make(map[string]FiberHandler),
		webServerPort:  port,
	}
}

func (wb *WebServer) Get(path string, handler FiberHandler) {
	wb.getHandlers[path] = handler
}

func (wb *WebServer) Post(path string, handler FiberHandler) {
	wb.postHandlers[path] = handler
}

func (wb *WebServer) Delete(path string, handler FiberHandler) {
	wb.deleteHandlers[path] = handler
}

func (wb *WebServer) Put(path string, handler FiberHandler) {
	wb.putHandlers[path] = handler
}

func (wb *WebServer) GetServer() *fiber.App {
	return wb.server
}

func (wb *WebServer) Start() error {
	wb.server.Use(logger.New())

	for path, hdl := range wb.getHandlers {
		wb.server.Get(path, hdl)
	}
	for path, hdl := range wb.postHandlers {
		wb.server.Post(path, hdl)
	}
	for path, hdl := range wb.deleteHandlers {
		wb.server.Delete(path, hdl)
	}
	for path, hdl := range wb.putHandlers {
		wb.server.Put(path, hdl)
	}

	return wb.server.Listen(fmt.Sprintf(":%s", wb.webServerPort))
}
