package utils

import "github.com/gofiber/fiber/v2"

func newResponse(c *fiber.Ctx, statusCode int, message string, data fiber.Map) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

func OkResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return newResponse(c, fiber.StatusOK, message, data)
}

func CreatedResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return newResponse(c, fiber.StatusCreated, message, data)
}

func AcceptedResponse(c *fiber.Ctx, message string) error {
	return newResponse(c, fiber.StatusAccepted, message, nil)
}

func BadRequestResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return newResponse(c, fiber.StatusBadRequest, message, data)
}

func NotFoundResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return newResponse(c, fiber.StatusNotFound, message, data)
}

func InternalServerErrorResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return newResponse(c, fiber.StatusInternalServerError, message, data)
}

func UnauthorizedResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return newResponse(c, fiber.StatusUnauthorized, message, data)
}

func ForbiddenResponse(c *fiber.Ctx, message string, data fiber.Map) error {
	return newResponse(c, fiber.StatusForbidden, message, data)
}
