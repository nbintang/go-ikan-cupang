package routes

import (
	"ikan-cupang/routes/auth"
	"ikan-cupang/routes/protected"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "This is api route",
		})
	})
	// Passing fiber.Router (api) to ProtectedRoutes
	protected.ProtectedRoutes(api)
	auth.AuthRoutes(api)
}
