package protected

import (
	protectedController "ikan-cupang/controllers/protected"
	"ikan-cupang/middlewares"
	"ikan-cupang/schemas"

	"github.com/gofiber/fiber/v2"
)

func FishRoutes(api fiber.Router) {
	apiGroup := api.Group("/fish")

	apiGroup.Get("/", protectedController.GetFishes)
	apiGroup.Get("/:id", protectedController.GetFishesById)
	apiGroup.Post("/",  middlewares.ValidateSchemaMiddleware(&schemas.CreateFishSchema{}), protectedController.CreateFish)
}
