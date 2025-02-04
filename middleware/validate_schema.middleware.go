package middleware

import (
	"ikan-cupang/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

type IError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateSchema(schema interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		newSchema := schema
		if err := c.BodyParser(newSchema); err != nil {
			return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid request body")
		}

		if err := Validator.Struct(newSchema); err != nil {
			var errors []IError
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, IError{
					Field: err.Field(),
					Tag:   err.Tag(),
					Value: err.Value().(string), 
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":  false,
				"messages": errors,
			})
		}

		c.Locals("validatedBody", newSchema)
		return c.Next()
	}
}
