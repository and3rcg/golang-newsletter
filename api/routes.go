package api

import "github.com/gofiber/fiber/v2"

func SetUpAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")

	apiGroup.Post("/send-email", SendEmail)
}
