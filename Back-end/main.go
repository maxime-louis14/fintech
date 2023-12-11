package main

import (
	"fintech/database"
	"fintech/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	database.ConnectDB()

	//routes
	routes.QuestionRoute(app) //add this

	app.Listen(":3001")
}
