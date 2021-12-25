package internal

import (
	router "goku/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
}

func MountRoutes(app *fiber.App) {
	app.Mount("/", router.BindDefaultRoutes())
	app.Mount("/api", router.BindApiRoutes())
}

func RunOnPort(port string) {
	app := fiber.New()

	SetupMiddleware(app)
	MountRoutes(app)

	log.Fatal(app.Listen(port))
}
