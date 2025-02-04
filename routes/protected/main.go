package protected

import (
	"github.com/gofiber/fiber/v2"
)

func ProtectedRoutes(api fiber.Router) {
	protected := api.Group("/protected")

	UserRoutes(protected)
}
