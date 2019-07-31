package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dnilosek/learning-web-dev/code/038-mongodb/006-update-model/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	collection *mongo.Collection
}

// Factory method
func NewUserController(c *mongo.Client) *UserController {
	collection := c.Database("go-web-dev-db").Collection("users")
	return &UserController{collection}
}

// Method for getting all users
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cur, err := uc.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get all restuls
	var results []models.User
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		results = append(results, elem)
	}

	// Marshal and send back to user
	uj, _ := json.Marshal(results)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// Method for getting user
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Get and find by id
	id := p.ByName("id")

	// Create hex from id and find it
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var u models.User
	err = uc.collection.FindOne(context.TODO(), bson.D{{"_id", idHex}}).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Marshal and send back to user
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// Method for creating user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	insertResult, err := uc.collection.InsertOne(context.TODO(), u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Inserted a single user: ", insertResult.InsertedID)

	// Send back
	id := struct {
		ID interface{} `json:id`
	}{
		ID: insertResult.InsertedID.(primitive.ObjectID),
	}
	idJ, _ := json.Marshal(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", idJ)
}

// Method for deleting user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Get and find by id
	id := p.ByName("id")

	// Create hex from id and find it
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	deleteResult, err := uc.collection.DeleteOne(context.TODO(), bson.D{{"_id", idHex}})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Delete Result:", deleteResult.DeletedCount)
	w.WriteHeader(http.StatusOK)
}
