package main

import (
	"fmt"
	"os"

	"github.com/ShardenduMishra22/DrStoneAPI/database"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var collection *mongo.Collection

func main() {
	fmt.Println("This is a DrStone API")

	// Creating a new Fiber instance for an Application
	app := fiber.New()

	// Setting up CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH,PUT,DELETE",
	}))

	// Loading .env file in Development mode
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}

	// Connecting to Database
	collection = database.ConnectToDataBase()

	// Test Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
	})

	// Listening to port
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "3000"
	}

	fmt.Println("Listening to port: " + Port)
	if err := app.Listen("0.0.0.0:" + Port); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
