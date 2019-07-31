package controllers

import "go.mongodb.org/mongo-driver/mongo"

type SessionController struct {
	client *mongo.Client
}

func NewSessionController(c *mongo.Client) *SessionController {
	return &SessionController{c}
}
