package migrations

import (
	"fmt"
	"ikan-cupang/config"
	"ikan-cupang/models"
	"log"
)

func DbMigration() {
	err := 	config.DB.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
		&models.VerificationToken{},
		&models.Fish{},
		&models.Category{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db migrated")
}
