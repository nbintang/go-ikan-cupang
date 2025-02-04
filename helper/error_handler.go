package helper

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx,status int, msg string,) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"message": msg,
	})
}
