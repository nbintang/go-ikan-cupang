package protected

import (
	"ikan-cupang/middlewares"
	"github.com/gofiber/fiber/v2"
)
func ProtectedRoutes(api fiber.Router) {
	protected := api.Group("/protected", middlewares.JwtBearerMiddleware)
	UserRoutes(protected)
	FishRoutes(protected)
}

