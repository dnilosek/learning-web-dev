package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat")
}

func main() {
	var d hotdog
	var c hotcat

	//	mux := http.NewServeMux()
	//	mux.Handle("/dog/", d) // localhost/dog/this/will/still/dog
	//	mux.Handle("/cat", c)  // localhost/cat/this/wont

	http.Handle("/dog/", d) // localhost/dog/this/will/still/dog
	http.Handle("/cat", c)  // localhost/cat/this/wont

	http.ListenAndServe(":8080", nil)
}
