package middlewares

import (
	"ikan-cupang/lib"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET"))

func JwtBearerMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Missing Authorization Header",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := lib.VerifyToken(tokenString)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid or expired token",
			"error":   err.Error(), // Tampilkan error parsing
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Token is not valid",
		})
	}

	c.Locals("user", token)
	return c.Next()
}
