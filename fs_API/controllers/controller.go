package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const collName = "watchlist"

var collection *mongo.Collection

// connecting with MONGODB

func init() {
	// client_options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to MDB
	client, err := mongo.Connect(context.TODO(), clientOption) // contexting the connection for use in later
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	connection = client.Database("GO_DB").Collection(collName)

	fmt.Println("Collection is ready!")
}
