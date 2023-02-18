package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "Url"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// Initilize with MongoBD
func init() {
	// It is a initilize method, run only one time

	// Client option
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	response, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully")

	collection = response.Database(dbName).Collection(colName)

	// Collection instance
	fmt.Println("Collection instance is ready")
}
