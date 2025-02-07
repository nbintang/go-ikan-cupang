package protected

import (
	"ikan-cupang/config"
	"ikan-cupang/daos"
	"ikan-cupang/dtos"
	"ikan-cupang/helper"
	"ikan-cupang/models"
	"ikan-cupang/schemas"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func GetFishes(c *fiber.Ctx) error {
	var fishes []*models.Fish
	db := config.DB

	if err := db.Debug().Preload("Category").Find(&fishes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helper.ErrorHandler(c, fiber.StatusNotFound, "Fish not found")
		}
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to get fishes")
	}
	var fishesResponse []dtos.Fish
	copier.CopyWithOption(&fishesResponse, &fishes, copier.Option{IgnoreEmpty: true})
	for i := range fishes {
		fishesResponse[i].Category = fishes[i].Category.Name
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    fishesResponse,
	})
}

func GetFishesById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid ID")
	}
	existedFish, err := daos.FindFishByID(uint(id))
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusNotFound, "Fish not found")
	}
	var fishesResponse dtos.Fish
	copier.CopyWithOption(&fishesResponse, &existedFish, copier.Option{IgnoreEmpty: true})

	fishesResponse.Category = existedFish.Category.Name

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    fishesResponse,
	})
}

func CreateFish(c *fiber.Ctx) error {
	body := c.Locals("validatedBody").(*schemas.FishSchema)
	db := config.DB
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid input Image")
	}
	imgURL, err := helper.UploadToCloudinary(fileHeader)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to upload image")
	}

	// Find or Create the Category
	var category models.Category
	if err := db.Where("name = ?", body.Category).FirstOrCreate(&category, models.Category{Name: body.Category}).Error; err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to create or find category")
	}

	newFish := models.Fish{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
		Image:       &imgURL,
		CategoryID:  category.ID,
	}
	if err := db.Create(&newFish).Error; err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to create fish")
	}
	db.Preload("Category").First(&newFish, newFish.ID)
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    newFish,
	})

}

func UpdateFish(c *fiber.Ctx) error {
	body := c.Locals("validatedBody").(*schemas.FishSchema)
	db := config.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid ID")
	}

	existedFish, err := daos.FindFishByID(uint(id))
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusNotFound, "Fish not found")
	}
	// Handle Image Upload
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid input Image")
	}

	imgURL, err := helper.UploadToCloudinary(fileHeader)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid input Image")
	}

	// Find or Create the Category
	var category models.Category
	if err := db.Where("name = ?", body.Category).FirstOrCreate(&category, models.Category{Name: body.Category}).Error; err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to create or find category")
	}

	// Update Fish Record
	if err := db.Model(&existedFish).Updates(models.Fish{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
		Image:       &imgURL,
		CategoryID:  category.ID, // Assign CategoryID
	}).Error; err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Failed to update fish")
	}

	// Reload fish data with category
	db.Preload("Category").First(&existedFish, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    existedFish,
	})
}

func DeleteFish(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	db := config.DB
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid ID")
	}

	existedFish, err := daos.FindFishByID(uint(id))
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusNotFound, "Fish not found")
	}
	if err := db.Where("id = ?", existedFish.ID).Delete(&existedFish).Error; err != nil {
		return helper.ErrorHandler(c, fiber.StatusInternalServerError, "Failed to delete fish")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    "Fish deleted",
	})
}
