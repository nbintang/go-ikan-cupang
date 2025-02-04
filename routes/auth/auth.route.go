package auth

import (

	"ikan-cupang/controllers/authentication"
	"ikan-cupang/middleware"
	"ikan-cupang/schemas"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoutes(api fiber.Router) {

	api.Post("/login", middleware.ValidateSchema(&schemas.LoginSchema{}), authentication.Login)
	api.Post("/verify-otp", middleware.ValidateSchema(&schemas.OTPSchema{}), authentication.VerifyOTP)
	api.Post("/resend-otp", middleware.ValidateSchema(&schemas.LoginSchema{}), authentication.ResendOTP)
	api.Post("/refresh-token", authentication.RefreshToken)
}
