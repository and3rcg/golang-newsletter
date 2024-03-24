package main

import (
	"newsletter-go/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: true,
		AppName:       "Newsletters in Go",
	})

	// setup routes
	api.SetUpAPIRoutes(app)

	app.Listen(":3000")
}
