package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupGetRoutes(app *fiber.App, coll *mongo.Collection) {
	collection = coll
	app.Get("/api/all", GetAllQuestions)
	app.Get("/api/question/:id", GetQuestionById)
	app.Get("/api/question/difficulty/:difficulty", GetQuestionByDifficulty)
}

func GetAllQuestions(c *fiber.Ctx) error {
	var qns []Question

	cursor, err := collection.Find(c.Context(), bson.M{})
	HandleError(err)
	defer cursor.Close(c.Context())

	for cursor.Next(c.Context()) {
		var qn Question
		err := cursor.Decode(&qn)
		HandleError(err)
		qns = append(qns, qn)
	}

	return c.Status(200).JSON(qns)
}

func GetQuestionById(c *fiber.Ctx) error {
	// code for finding question by id
	return nil
}

func GetQuestionByDifficulty(c *fiber.Ctx) error {
	// code for finding question by difficulty
	return nil
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
	fmt.Println("There was no error!!")
}
