package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dnilosek/learning-web-dev/code/038-mongodb/004-controllers/models"
	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

// Factory method
func NewUserController() *UserController {
	return &UserController{}
}

// Method for getting user
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		ID:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// Method for creating user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.ID = "007"

	// Send back
	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// Method for deleting user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Okay
	w.WriteHeader(http.StatusOK)
}
