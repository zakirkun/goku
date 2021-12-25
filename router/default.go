package router

import (
	"github.com/gofiber/fiber/v2"
)

func BindDefaultRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("APP")
	})

	return app
}
