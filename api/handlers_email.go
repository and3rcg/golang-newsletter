package api

import (
	"newsletter-go/internal"
	"newsletter-go/utils"

	"github.com/gofiber/fiber/v2"
)

// wip
func SendEmailHandler(c *fiber.Ctx, repo *internal.Repository) error {
	return utils.OkResponse(c, "Sending e-mail...", nil)
}
