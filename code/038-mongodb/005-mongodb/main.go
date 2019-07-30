package main

import (
	"net/http"

	"github.com/dnilosek/learning-web-dev/code/038-mongodb/005-mongodb/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uc := controllers.NewUserController(getClient())

	router := httprouter.New()
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", router)
}

func getClient() *mongo.Client {
	// Connect to local mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	// Check connetion
	if err != nil {
		panic(err)
	}
	return client
}
