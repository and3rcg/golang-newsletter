package api

import (
	"newsletter-go/internal"
	"newsletter-go/models"
	"newsletter-go/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateNewsletterHandler creates a newsletter instance
func CreateNewsletterHandler(c *fiber.Ctx, r *internal.Repository) error {
	var request models.Newsletter

	err := c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "Failed to parse request body")
	}

	validationErrs := r.Validator.Struct(request)
	if validationErrs != nil {
		return utils.BadRequestResponse(c, validationErrs.Error())
	}

	err = CreateNewsletterOperation(r, &request)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.CreatedResponse(c, "Newsletter created successfully", fiber.Map{
		"object": request,
	})
}

// GetNewsletterHandler retrieves a single newsletter instance
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

// GetNewsletterListHandler retrieves a list of all newsletters
func GetNewsletterListHandler(c *fiber.Ctx, r *internal.Repository) error {
	newsletterList, err := GetAllNewslettersOperation(r)

	if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.OkResponse(c, "Newsletter list retrieved successfully", fiber.Map{
		"object_list": newsletterList,
	})
}

// UpdateNewsletterHandler updates a newsletter instance
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

	validationErrs := r.Validator.Struct(request)
	if validationErrs != nil {
		return utils.BadRequestResponse(c, validationErrs.Error())
	}

	err = UpdateNewsletterOperation(r, &request, id)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "Newsletter not found")
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.OkResponse(c, "Newsletter updated successfully", nil)
}

// DeleteNewsletterHandler deletes a newsletter instance
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

// SubscribeToNewsletterHandler adds an email address to the specified newsletter. Duplicate email addresses are not allowed.
func SubscribeToNewsletterHandler(c *fiber.Ctx, r *internal.Repository) error {
	var request models.NewsletterUser

	err := c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "Failed to parse request body")
	}

	validationErrs := r.Validator.Struct(request)
	if validationErrs != nil {
		return utils.BadRequestResponse(c, validationErrs.Error())
	}

	err = SubscribeToNewsletterOperation(r, &request)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "Newsletter not found")
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.OkResponse(c, "E-mail added successfully", nil)
}

// UnsubscribeFromNewsletterHandler removes an e-mail address from the specified newsletter. Duplicates are not allowed.
func UnsubscribeFromNewsletterHandler(c *fiber.Ctx, r *internal.Repository) error {
	email := c.Query("email", "")
	newsletter_id := c.QueryInt("newsletter_id", -1)

	if email == "" || newsletter_id == -1 {
		return utils.BadRequestResponse(c, "Invalid request")
	}

	err := UnsubscribeFromNewsletterOperation(r, email, uint(newsletter_id))
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "Newsletter not found")
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	return utils.OkResponse(c, "E-mail removed successfully", nil)
}
