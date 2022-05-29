package controller

import (
	"context"
	"fmt"
	"log"

	model "github.com/Sam44323/fs_API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// MONGODB helpers

// insert a record

func insertOneMovie(movie model.Netflix) {
	data, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a movie: ", data.InsertedID)
}

// update a record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}                       // creating the query for filter
	update := bson.M{"$set": bson.M{"watched": true}} // creating the query for updates

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie updated! ", res.ModifiedCount)
}

// delete one record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie deleted! ", res.DeletedCount)
}

// delete multiple_records

func deleteManyMovies() {
	filter := primitive.M{{}} // it means delete everything

	res, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie deleted! ", res.DeletedCount)
}
