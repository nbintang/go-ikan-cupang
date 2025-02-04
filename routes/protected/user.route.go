package protected

import (
	"ikan-cupang/controllers/protected"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	apiGroup := api.Group("/users")
	apiGroup.Get("/", protected.GetUsers) 
	apiGroup.Get("/:id", protected.GetUserByID)
}

