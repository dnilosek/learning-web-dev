package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
	Last  string
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	p1 := person{
		First: "James",
		Last:  "Bond",
	}
	// Header not required, good practice
	res.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(res)
	err := enc.Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
