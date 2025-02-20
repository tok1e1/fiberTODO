package routes

import (
	"fiberTODO/cmd/database"
	"fiberTODO/server"
	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App, db database.DB) {
	api := app.Group("/tasks")

	api.Post("/", func(c *fiber.Ctx) error {
		return server.CreateTasks(c, db)
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return server.GetTasks(c, db)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return server.UpdateTask(c, db)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return server.DeleteTask(c, db)
	})
}
