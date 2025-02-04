package daos

import (
	"ikan-cupang/config"
	"ikan-cupang/models"
)

func FindUserByEmail(email string, selection ...string) (*models.User, error) {
	var user models.User
	db := config.DB
	err := db.Select(selection).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}


func FindUserByID(id uint, selection ...string) (*models.User, error) {
	var user models.User
	db := config.DB
	err := db.Select(selection).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
