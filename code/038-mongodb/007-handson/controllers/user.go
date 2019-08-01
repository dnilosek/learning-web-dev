package controllers

import (
	"context"
	"fmt"

	"github.com/dnilosek/learning-web-dev/code/038-mongodb/007-handson/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	collection *mongo.Collection
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c.Database("go-web-dev-db").Collection("users")}
}

// Check if username exists
func (uc *UserController) UsernameExists(username string) bool {

	err := uc.collection.FindOne(context.TODO(), bson.D{{"username", username}})
	if err != nil {
		return false
	}
	return true
}

// Get user from db
func (uc *UserController) GetUser(username string) (models.User, error) {
	var u models.User
	err := uc.collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

// Create a user, return their ID
func (uc *UserController) CreateUser(u models.User) (string, error) {
	var uId string

	//Marshal to JSON and insert
	insertResult, err := uc.collection.InsertOne(context.TODO(), u)
	if err != nil {
		return uId, err
	}

	uId = fmt.Sprint(insertResult.InsertedID.(primitive.ObjectID).Hex())
	return uId, nil
}
