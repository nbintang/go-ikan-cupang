package auth

import (
	"ikan-cupang/controllers/authentication"

	"ikan-cupang/middlewares"
	"ikan-cupang/schemas"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoutes(api fiber.Router) {

	api.Post("/login", middlewares.ValidateSchemaMiddleware(&schemas.LoginSchema{}), authentication.Login)
	api.Post("/verify-otp", middlewares.ValidateSchemaMiddleware(&schemas.OTPSchema{}), authentication.VerifyOTP)
	api.Post("/resend-otp", middlewares.ValidateSchemaMiddleware(&schemas.LoginSchema{}), authentication.ResendOTP)
	api.Post("/refresh-token", authentication.RefreshToken)
}
