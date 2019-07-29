package main

import (
	"encoding/json"
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
	enc := json.NewEncoder(res)
	enc.Encode(p1)
}
