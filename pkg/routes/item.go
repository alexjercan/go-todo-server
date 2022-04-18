package routes

import (
	"context"
	"time"

	"github.com/alexjercan/go-todo-server/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func ItemList(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	items, err := database.GetItems(ctx)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(items)
}
