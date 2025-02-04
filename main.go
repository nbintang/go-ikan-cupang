package main

import (
	"fmt"
	"ikan-cupang/config"
	"ikan-cupang/config/migrations"
	"ikan-cupang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET, POST, PUT, DELETE",
		}),
	)
	config.DbInit()
	migrations.DbMigration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "hello world",
		})
	})

	routes.ApiRoutes(app)

	fmt.Println("Server running at http://localhost:3001")
	if err := app.Listen(":3001"); err != nil {
		log.Fatal(err)
	}
}
