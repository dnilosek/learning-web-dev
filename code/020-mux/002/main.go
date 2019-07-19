package main

import (
	"io"
	"net/http"
)

func hotdog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func hotcat(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat")
}

func main() {

	http.HandleFunc("/dog/", hotdog) // localhost/dog/this/will/still/dog
	http.HandleFunc("/cat", hotcat)  // localhost/cat/this/wont

	http.ListenAndServe(":8080", nil)
}
