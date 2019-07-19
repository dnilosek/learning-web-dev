package main

import (
	"io"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "root")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "me")
}

func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
