package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	if ve, ok := err.(validator.ValidationErrors); ok {
		code = fiber.StatusBadRequest
		return c.Status(code).JSON(fiber.Map{
			"error":   true,
			"code":    code,
			"message": "Validation failed",
			"details": ParseValidationErrors(ve),
		})
	}

	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"code":    code,
		"message": message,
	})
}

