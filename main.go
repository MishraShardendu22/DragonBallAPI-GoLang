package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("This is a DrStone API")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH,PUT,DELETE",
	}))

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
	})
	
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "3000" 
	}

	fmt.Println("Listening to port: " + Port)
	if err := app.Listen("0.0.0.0:" + Port); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
