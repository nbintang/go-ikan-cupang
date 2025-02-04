package main

import (
	"fmt"
	"ikan-cupang/config"
	"ikan-cupang/models"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
}

type Order struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func SeedDB(db *gorm.DB) {
	description := "Fresh tuna fish"
	imgUrl := "https://example.com/tuna.jpg"
	fish := []models.Fish{
		{
			Name:        "Tuna",
			Description: &description,
			Price:       10.99,
			Stock:       100,
			Image:       &imgUrl,
			CategoryID:  1,
		},
		{
			Name:        "Tuna 2",
			Description: &description,
			Price:       20.54,
			Stock:       20,
			Image:       &imgUrl,
			CategoryID:  1,
		},
	}

	for i := range fish {
		if err := db.Create(&fish[i]).Error; err != nil {
			fmt.Println("Error seeding data:", err)
			return
		}
	}



	fmt.Println("Database seeded successfully")
}

func main() {
    if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	config.DbInit()
	SeedDB(config.DB)
}
