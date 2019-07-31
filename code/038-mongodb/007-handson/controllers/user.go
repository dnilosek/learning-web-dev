package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

// Get user from db
func (uc *UserController) GetUser(userId string) error {

	return nil
}
