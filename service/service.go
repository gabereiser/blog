package service

import (
	"github.com/gabereiser/blog/controllers"
	"github.com/gofiber/fiber/v2"
)

// Service is the interface that wraps the basic methods for a service.
type Service interface {
	// Start initializes the service and starts the server.
	Start() error
	// Stop gracefully stops the service.
	Stop() error
	// GetFiberApp returns the fiber app instance.
	GetFiberApp() *fiber.App
}

type WebService struct {
	app *fiber.App
}

func NewWebService() *WebService {
	app := fiber.New(fiber.Config{
		AppName:           "Blog",
		EnablePrintRoutes: true,
		ETag:              true,
	})
	return &WebService{app: app}
}

func (ws *WebService) Start() error {
	// Start the Fiber app
	return ws.app.Listen(":8080")
}

func (ws *WebService) Stop() error {
	// Gracefully stop the Fiber app
	return ws.app.Shutdown()
}

func (ws *WebService) GetFiberApp() *fiber.App {
	return ws.app
}

func (ws *WebService) RegisterRoutes() {
	// Register your routes here
	controllers.NewHomeController().RegisterRoutes(ws.app)
	controllers.NewAuthController().
		RegisterAnonymousRoutes(ws.app).
		RegisterRoutes(ws.app)

	controllers.NewBlogController().RegisterRoutes(ws.app)
}
