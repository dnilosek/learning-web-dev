//usr/bin/go run $0 $@ ; exit
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello from docker")
}
