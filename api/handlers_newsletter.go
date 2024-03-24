package api

import (
	"newsletter-go/internal"
	"newsletter-go/models"
	"newsletter-go/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateNewsletterHandler(c *fiber.Ctx, r *internal.Repository) error {
	var request models.Newsletter

	err := c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "Failed to parse request body")
	}

	err = CreateNewsletterOperation(r, &request)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.CreatedResponse(c, "Newsletter created successfully", fiber.Map{
		"object": request,
	})
}

func GetNewsletterHandler(c *fiber.Ctx, r *internal.Repository) error {
	idString := c.Params("id", "")

	id, err := strconv.Atoi(idString)

	if err != nil {
		return utils.BadRequestResponse(c, "Invalid newsletter ID")
	}

	newsletterObj, err := GetNewsletterByIDOperation(r, id)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "Newsletter not found")
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.OkResponse(c, "Newsletter retrieved successfully", fiber.Map{
		"object": newsletterObj,
	})
}

func GetNewsletterListHandler(c *fiber.Ctx, r *internal.Repository) error {
	newsletterList, err := GetAllNewslettersOperation(r)

	if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.OkResponse(c, "Newsletter list retrieved successfully", fiber.Map{
		"object_list": newsletterList,
	})
}

func UpdateNewsletterHandler(c *fiber.Ctx, r *internal.Repository) error {
	var request models.Newsletter
	idString := c.Params("id", "")

	id, err := strconv.Atoi(idString)

	if err != nil {
		return utils.BadRequestResponse(c, "Invalid newsletter ID")
	}

	err = c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "Failed to parse request body")
	}

	err = UpdateNewsletterOperation(r, &request, id)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "Newsletter not found")
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.OkResponse(c, "Newsletter updated successfully", nil)
}

func DeleteNewsletterHandler(c *fiber.Ctx, r *internal.Repository) error {
	idString := c.Params("id", "")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid newsletter ID")
	}

	err = DeleteNewsletterOperation(r, id)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "Newsletter not found")
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.OkResponse(c, "Newsletter deleted successfully", nil)
}
