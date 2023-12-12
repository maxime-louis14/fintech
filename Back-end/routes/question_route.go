package routes

import (
	"fintech/controllers"
	"fintech/middleware"
	"github.com/gofiber/fiber/v2"
)

func QuestionRoute(app *fiber.App) {
	// Utilisation du middleware CORS
	app.Use(middleware.CorsMiddleware())

	// Route pour la méthode GET
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Route pour la méthode POST vers /form
	app.Post("/form", controllers.PostForm)
	app.Post("/question", controllers.PostQuestionnaires)
}
