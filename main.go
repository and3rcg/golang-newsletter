package main

import (
	"log"
	"newsletter-go/api"
	"newsletter-go/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: true,
		AppName:       "Newsletters in Go",
	})

	repo, err := internal.StartRepository()
	if err != nil {
		log.Fatal(err)
	}

	// setup routes
	api.SetUpAPIRoutes(app, repo)

	app.Listen(":3000")
}
