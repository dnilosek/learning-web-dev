package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", ping)

	http.ListenAndServe(":80", nil)
}

func ping(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "OK")
}
