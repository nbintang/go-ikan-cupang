package auth

import "github.com/gofiber/fiber/v2"

func AuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	AuthenticationRoutes(auth)

}