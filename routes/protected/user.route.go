package protected

import (
	"ikan-cupang/controllers/protected"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	api.Get("/users", protected.GetUsers) 
	api.Get("/users/:id", protected.GetUserByID)
}

