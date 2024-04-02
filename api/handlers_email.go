package api

import (
	"newsletter-go/internal"
	"newsletter-go/models"
	"newsletter-go/tasks"
	"newsletter-go/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func SendNewsletterEmailsHandler(c *fiber.Ctx, r *internal.Repository) error {
	var request models.EmailContent
	err := c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "Failed to parse request body")
	}

	validationErrs := r.Validator.Struct(request)
	if validationErrs != nil {
		return utils.BadRequestResponse(c, validationErrs.Error())
	}

	newsletterIDStr := c.Params("id", "")
	newsletterID, err := strconv.Atoi(newsletterIDStr)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid newsletter ID")
	}

	err = c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "Failed to parse request body")
	}

	newsletterObj, err := GetNewsletterByIDOperation(r, newsletterID)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "Newsletter not found")
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	task, err := tasks.NewTaskSendNewsletterEmails(c, *newsletterObj, request, "emails/email_default.html")
	if err != nil {
		return utils.InternalServerErrorResponse(c, err.Error())
	}

	client := tasks.GetWorkerClient()
	client.Enqueue(task, asynq.Queue("critical"))

	return utils.OkResponse(c, "Sending e-mail...", fiber.Map{
		"data":          request,
		"newsletter_id": newsletterID,
	})
}
