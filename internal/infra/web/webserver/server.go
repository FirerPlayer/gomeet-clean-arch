package webserver

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberHandler func(c *fiber.Ctx) error

type WebServer struct {
	app            *fiber.App
	getHandlers    map[string]FiberHandler
	postHandlers   map[string]FiberHandler
	deleteHandlers map[string]FiberHandler
	putHandlers    map[string]FiberHandler
	webServerPort  string
}

func NewWebServer(port, appName string) *WebServer {
	return &WebServer{
		app: fiber.New(fiber.Config{
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

func (wb *WebServer) GetApp() *fiber.App {
	return wb.app
}

func (wb *WebServer) Start() error {
	wb.app.Use(logger.New())
	wb.app.Use(cors.New())
	apiR := wb.app.Group("/api")

	for path, hdl := range wb.getHandlers {
		apiR.Get(path, hdl)
	}
	for path, hdl := range wb.postHandlers {
		apiR.Post(path, hdl)
	}
	for path, hdl := range wb.deleteHandlers {
		apiR.Delete(path, hdl)
	}
	for path, hdl := range wb.putHandlers {
		apiR.Put(path, hdl)
	}

	return wb.app.Listen(fmt.Sprintf(":%s", wb.webServerPort))
}
