package routes

import (
	"github.com/gofiber/fiber/v2"
)

func QuestionRoute(app *fiber.App) {
	// Simple GET handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
