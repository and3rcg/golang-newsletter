package api

import (
	"newsletter-go/internal"

	"github.com/gofiber/fiber/v2"
)

func SetUpAPIRoutes(app *fiber.App, repo *internal.Repository) {
	newsletterGroup := app.Group("/api/newsletter")

	newsletterGroup.Post("/", func(c *fiber.Ctx) error {
		return CreateNewsletterHandler(c, repo)
	})

	newsletterGroup.Get("/", func(c *fiber.Ctx) error {
		return GetNewsletterListHandler(c, repo)
	})

	newsletterGroup.Get("/:id", func(c *fiber.Ctx) error {
		return GetNewsletterHandler(c, repo)
	})

	newsletterGroup.Patch("/:id", func(c *fiber.Ctx) error {
		return UpdateNewsletterHandler(c, repo)
	})

	newsletterGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return DeleteNewsletterHandler(c, repo)
	})

	emailGroup := app.Group("/api/email")

	emailGroup.Post("/send", func(c *fiber.Ctx) error {
		return SendEmailHandler(c, repo)
	})
}
