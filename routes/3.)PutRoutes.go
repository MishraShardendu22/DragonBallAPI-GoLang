package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupPutRoutes(app *fiber.App, coll *mongo.Collection){
	collection = coll
	// app.Put("/api/question/:QuestionNumber", UpdateQuestion)
}

// func UpdateQuestion(c *fiber.Ctx)error {
// 	id := c.Params("id")
	
// }