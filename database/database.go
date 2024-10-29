package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDataBase() *mongo.Collection {
	fmt.Println("Connecting To The Database !!")

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set. Set the Database URI to connect.")
	}
	clientOption := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Ping the database to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	// Return the connection to the desired collection
	return client.Database("DrStone").Collection("DrStone")
}

// Steps To COnnect Database - 

// 1.) Make a ConnectToDatabase function with return type "*mongo.Collection"
// 2.) Get the MongoURI (from env file) 
// 3.) Make a clientOption
// 4.) Make a client and Connect to the database 
// 5.) Ping the database to verify connection
// 6.) Return the connection to the desired collection and database 