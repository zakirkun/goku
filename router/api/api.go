package api

import (
	"github.com/gofiber/fiber/v2"
)

func BindApiRoutes() *fiber.App {
	api := fiber.New()

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API")
	})

	return api
}
