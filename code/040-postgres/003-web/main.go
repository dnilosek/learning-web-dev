package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Book struct {
	ISBN   string  `json:isbn`
	Title  string  `json:title`
	Author string  `json:author`
	Price  float32 `json:price`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://gopher:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/books", index)
	http.HandleFunc("/books/show", isbn)
	http.ListenAndServe(":8080", nil)
}

func isbn(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	row := db.QueryRow(`SELECT * FROM books WHERE isbn = $1;`, isbn)

	bk := Book{}
	err := row.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price) // order matters
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(res, req)
		return
	case err != nil:
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	bkJson, err := json.Marshal(bk)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(res, "%s\n", bkJson)
}
func index(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query(`SELECT * FROM books;`)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price) // order matters
		if err != nil {
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		bks = append(bks, bk)
	}

	bksJson, err := json.Marshal(bks)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(res, "%s\n", bksJson)
}
