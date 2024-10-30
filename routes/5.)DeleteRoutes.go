package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupDeleteRoutes(app *fiber.App, coll *mongo.Collection){
	collection = coll
	app.Delete("/api/question/:QuestionNumber", DeleteQuestion)
	app.Delete("/api/delete_all", DeleteAll)
}

func DeleteQuestion(c *fiber.Ctx) error {
	
}

func DeleteAll(c *fiber.Ctx) error {

}