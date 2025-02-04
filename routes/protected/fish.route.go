package protected

import (
	"ikan-cupang/controllers/protected"

	"github.com/gofiber/fiber/v2"
)

func FishRoutes(api fiber.Router) {
 apiGroup := api.Group("/fish")

 apiGroup.Get("/", protected.GetFishes)
}