package protected

import (
	"ikan-cupang/config"
	"ikan-cupang/dtos"
	"ikan-cupang/helper"
	"ikan-cupang/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func GetFishes(c *fiber.Ctx) error {
	var fish []*models.Fish
	db := config.DB

	if err := db.Debug().Find(&fish).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.ErrorHandler(c, fiber.StatusNotFound, "Fish not found")
		}
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to get fish")
	}

	var fishesResponse []dtos.Fish

	if err := copier.Copy(&fishesResponse, &fish); err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to copy fish")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    fishesResponse,
	})
}
