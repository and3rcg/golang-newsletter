package api

import (
	"newsletter-go/models"
	"newsletter-go/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateNewsletter(c *fiber.Ctx) error {
	return utils.CreatedResponse(c, "Newsletter created successfully", fiber.Map{})
}

func GetNewsletter(c *fiber.Ctx) error {
	var newsletterObj models.Newsletter
	return utils.OkResponse(c, "Newsletter retrieved successfully", fiber.Map{
		"object": newsletterObj,
	})
}

func GetNewsletterList(c *fiber.Ctx) error {
	var newsletterList []models.Newsletter
	return utils.OkResponse(c, "Newsletter list retrieved successfully", fiber.Map{
		"object_list": newsletterList,
	})
}

func UpdateNewsletter(c *fiber.Ctx) error {
	return utils.OkResponse(c, "Newsletter updated successfully", nil)
}

func DeleteNewsletter(c *fiber.Ctx) error {
	return utils.OkResponse(c, "Newsletter deleted successfully", nil)
}
