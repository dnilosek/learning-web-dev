package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//clientOptions.SetAuth(options.Credential{AuthMechanism: "SCRAM-SHA-1", Username: "bond", Password: "test"})
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}
	DB = client.Database("bookstore").Collection("books")
	fmt.Println("Connected to MongoDB!")
}
