package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupPostRoutes(app *fiber.App, coll *mongo.Collection) {
	collection = coll
	app.Post("/api/question", AddQuestion)
}

func AddQuestion(c *fiber.Ctx) error {
	var qn Question
	err := c.BodyParser(&qn)
	HandleError(err)

	if qn.Rating < 0 {
		qn.Rating = 50
	}

	if qn.Difficulty == "" {
		qn.Difficulty = "Medium"
	}

	_, err = collection.InsertOne(context.Background(), qn)
	HandleError(err)

	return c.Status(200).JSON(fiber.Map{"message" : qn})
}
