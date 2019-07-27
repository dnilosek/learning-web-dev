package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {
	connStr := "postgres://david.nilosek:Nyrod2035@localhost/test?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Success")
	check(err)
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
