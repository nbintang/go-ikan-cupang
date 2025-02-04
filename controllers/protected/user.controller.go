package protected

import (
	"ikan-cupang/config"
	"ikan-cupang/dtos"
	"ikan-cupang/helper"
	"ikan-cupang/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetUsers(c *fiber.Ctx) error {
	var users []*models.User
	db := config.DB
	db.Debug().Where("role = ?", "USER").Find(&users)

	var userResponse []dtos.User
	copier.Copy(&userResponse, &users)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    userResponse,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	var user *models.User
	idParams := c.Params("id")
	db := config.DB
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusBadRequest, "Invalid ID")
	}

	err = db.Debug().Where("role = ?", "USER").First(&user, "id = ?", id).Error
	if err != nil {
		return helper.ErrorHandler(c, fiber.StatusNotFound, "User not found")
	}

	var userResponse dtos.User
	copier.Copy(&userResponse, &user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    userResponse,
	})
}



