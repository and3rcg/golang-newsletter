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

	newsletterGroup.Post("/subscribe", func(c *fiber.Ctx) error {
		return SubscribeToNewsletterHandler(c, repo)
	})

	newsletterGroup.Post("/unsubscribe", func(c *fiber.Ctx) error {
		return UnsubscribeFromNewsletterHandler(c, repo)
	})

	newsletterGroup.Post("/:id/send-emails", func(c *fiber.Ctx) error {
		return SendNewsletterEmailsHandler(c, repo)
	})
}
