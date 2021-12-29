package web

import (
	"github.com/gofiber/fiber/v2"
)

func BindWebRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("APP")
	})

	return app
}
