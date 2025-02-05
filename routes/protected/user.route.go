package protected

import (
	protectedController"ikan-cupang/controllers/protected"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	apiGroup := api.Group("/users")
	apiGroup.Get("/", protectedController.GetUsers) 
	apiGroup.Get("/:id", protectedController.GetUserByID)
}

