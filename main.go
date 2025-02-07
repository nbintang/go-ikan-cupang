package main

import (
	"fmt"
	"ikan-cupang/config"
	"ikan-cupang/config/migrations"
	"ikan-cupang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New();
	app.Use(
		cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins: "http://localhost:3000",
			AllowMethods: "GET, POST, PUT, DELETE, PATCH",
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
	port := 3001
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}
