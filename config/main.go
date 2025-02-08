package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() {
	var err error
	dsn := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error", err)
	}
	fmt.Println("db connected")
}
