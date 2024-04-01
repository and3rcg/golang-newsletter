package api

import (
	"newsletter-go/internal"
	"newsletter-go/tasks"
	"newsletter-go/utils"

	"github.com/gofiber/fiber/v2"
)

func SetUpAPIRoutes(app *fiber.App, repo *internal.Repository) {
	newsletterGroup := app.Group("/api/newsletter")

	newsletterGroup.Get("/demo-task", func(c *fiber.Ctx) error {
		client := tasks.GetWorkerClient()
		task := tasks.NewTaskDemo(10)
		client.Enqueue(task)
		return utils.OkResponse(c, "task started", nil)
	})

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
		return UnsubscribeFromNewsletterAPIHandler(c, repo)
	})

	newsletterGroup.Post("/:id/send-emails", func(c *fiber.Ctx) error {
		return SendNewsletterEmailsHandler(c, repo)
	})

	app.Get("/unsubscribe", func(c *fiber.Ctx) error {
		return UnsubscribeFromNewsletterHandler(c, repo)
	})
}
