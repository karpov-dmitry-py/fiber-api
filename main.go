package main

import (
	"github.com/karpov-dmitry-py/fiber-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/karpov-dmitry-py/fiber-api/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("API welcome end point")
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	if _, err := database.ConnectDb(); err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listen(":5000"))
}
