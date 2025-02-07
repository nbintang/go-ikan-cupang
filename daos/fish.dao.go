package daos

import (
	"ikan-cupang/config"
	"ikan-cupang/models"
)

func FindFishByID(id uint, selection ...string) (*models.Fish, error) {
	var fish models.Fish
	db := config.DB
	err := db.Debug().Preload("Category").Select(selection).Where("id = ?", id).First(&fish).Error
	if err != nil {
		return nil, err
	}
	return &fish, nil
}