package exception

import (
	"errors"
	"go-resto/app/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	var ve validator.ValidationErrors
	if errors.As(err, &e) {
		code = e.Code
	} else if errors.As(err, &ve) {
		code = fiber.StatusBadRequest
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Return status code with error message
	return helpers.ErrResponse(c, err.Error(), code)
}
