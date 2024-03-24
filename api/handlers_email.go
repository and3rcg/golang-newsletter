package api

import (
	"newsletter-go/utils"

	"github.com/gofiber/fiber/v2"
)

func SendEmail(c *fiber.Ctx) error {
	return utils.OkResponse(c, "Sending e-mail...", nil)
}
