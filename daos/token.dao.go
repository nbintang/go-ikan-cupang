package daos

import (
	"ikan-cupang/config"
	"ikan-cupang/models"
)

func DeleteAllTokenByUserID(id uint) (*models.VerificationToken, error) {
	var token models.VerificationToken
	db := config.DB
	err := db.Where("user_id = ?", id).Delete(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func FindExistedTokenByUserID(id uint) (*models.VerificationToken, error) {
	var token models.VerificationToken
	db := config.DB
	err := db.Where("user_id = ?", id).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}
