package main

import (
	"os"

	"github.com/alexjercan/go-todo-server/pkg/database"
	"github.com/alexjercan/go-todo-server/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	connection := database.Connection{
		Address:  os.Getenv("DB_ADDRESS"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
	database.Connect(connection)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/items", routes.ItemList)
	// app.Post("/api/items", routes.ItemCreate)
	// app.Put("/api/items", routes.UpdateItem)
	// app.Delete("/api/items", routes.DeleteItem)

	app.Listen(":3000")
}
