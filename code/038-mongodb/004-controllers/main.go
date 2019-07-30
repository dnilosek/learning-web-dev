package main

import (
	"net/http"

	"github.com/dnilosek/learning-web-dev/code/038-mongodb/004-controllers/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	uc := controllers.NewUserController()

	router := httprouter.New()
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", router)
}
