package controllers

import "github.com/gofiber/fiber/v2"

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}
func (hc *HomeController) RegisterRoutes(app *fiber.App) {
	app.Get("/", home)
	app.Get("/status", status)
}

func home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func status(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Server is running",
	})
}
