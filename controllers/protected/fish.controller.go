package protected

import (
	"ikan-cupang/config"
	"ikan-cupang/dtos"
	"ikan-cupang/helper"
	"ikan-cupang/models"
	"ikan-cupang/schemas"

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

func GetFishesById(c *fiber.Ctx) error {
	var fish *models.Fish
	db := config.DB
	id := c.Params("id")
	if err := db.Debug().First(&fish, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.ErrorHandler(c, fiber.StatusNotFound, "Fish not found")
		}

		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to get fish")
	}

	var fishesResponse dtos.Fish

	if err := copier.Copy(&fishesResponse, &fish); err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to copy fish")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    fishesResponse,
	})
}

func CreateFish(c *fiber.Ctx) error {
	body := c.Locals("validatedBody").(*schemas.CreateFishSchema)
	db := config.DB
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid input Image")
	}
	imgURL, err := helper.UploadToCloudinary(fileHeader)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to upload image")
	}

	newFish := models.Fish{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
		Image:       &imgURL,
		Category: models.Category{
			Name: body.Category,
		},
	}
	if err := db.Create(&newFish).Error; err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to create fish")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    newFish,
	})

}
