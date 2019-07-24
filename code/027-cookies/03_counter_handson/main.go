package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/count", count)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	// increment counter
	cookie := incCookie(req)
	http.SetCookie(res, cookie)
}

func count(res http.ResponseWriter, req *http.Request) {
	// count
	cookie := incCookie(req)
	http.SetCookie(res, cookie)

	fmt.Fprintf(res, "You've visited %v times", cookie.Value)
}

func incCookie(req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("counter")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "counter",
			Value: "0",
			Path:  "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	return cookie
}
