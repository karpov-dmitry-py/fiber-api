package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/karpov-dmitry-py/fiber-api/database"
	"github.com/karpov-dmitry-py/fiber-api/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
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

	log.Fatal(app.Listen(":8000"))
}
