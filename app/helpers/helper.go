package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate *validator.Validate = validator.New()

func OkResponse(c *fiber.Ctx, message string, data any) error {
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func ErrResponse(c *fiber.Ctx, message string, code int) error {
	return c.Status(code).JSON(map[string]interface{}{
		"success": false,
		"message": message,
	})
}
