package main

import (
	"log"
	"newsletter-go/api"
	"newsletter-go/internal"
	"newsletter-go/tasks"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	engine := html.New("./templates/views", ".html")
	app := fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: true,
		AppName:       "Newsletters in Go",
		Views:         engine,
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repo, err := internal.StartRepository()
	if err != nil {
		log.Fatal(err)
	}

	// starting up the worker client
	tasks.InitWorkerClient()
	defer tasks.CloseWorkerClient()

	// setup routes
	api.SetUpAPIRoutes(app, repo)

	app.Listen(":3000")
}
