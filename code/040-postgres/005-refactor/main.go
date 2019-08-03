package main

import (
	"net/http"

	"github.com/dnilosek/learning-web-dev/code/040-postgres/005-refactor/books"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", books.BooksIndex)
	http.HandleFunc("/books/show", books.BooksShow)
	http.HandleFunc("/books/create", books.BooksCreateForm)
	http.HandleFunc("/books/create/process", books.BooksCreateProcess)
	http.HandleFunc("/books/update", books.BooksUpdateForm)
	http.HandleFunc("/books/update/process", books.BooksUpdateProcess)
	http.HandleFunc("/books/delete", books.BooksDelete)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/books", http.StatusSeeOther)
}
