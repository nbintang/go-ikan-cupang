package main

import (
	"fmt"
	"ikan-cupang/config"
	"ikan-cupang/models"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func DbMigration() {
	err := config.DB.AutoMigrate(
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

	// show all tables
	// Show all models with their fields and types
	showModelDetails(config.DB)
}

// run migrations with go run config/migrations/init.go
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	config.DbInit()
	DbMigration()
}

func showModelDetails(db *gorm.DB) {
	var tables []string
	db.Raw("SHOW TABLES").Scan(&tables)

	fmt.Println("\n📌 Migrated Models and Fields:")
	for _, table := range tables {
		fmt.Printf("\n🔹 Table: %s\n", table)

		// Get column details
		columns, err := db.Migrator().ColumnTypes(table)
		if err != nil {
			fmt.Printf("Error fetching columns for %s: %v\n", table, err)
			continue
		}

		// Print column names and their data types
		for _, column := range columns {
			fmt.Printf("   - %s (%s)\n", column.Name(), column.DatabaseTypeName())
		}
	}
}
