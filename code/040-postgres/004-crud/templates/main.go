package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Book struct {
	ISBN   string  `json:isbn`
	Title  string  `json:title`
	Author string  `json:author`
	Price  float32 `json:price`
}

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://gopher:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create", booksCreateForm)
	http.HandleFunc("/books/create/process", booksCreateProcess)
	http.HandleFunc("/books/update", booksUpdateForm)
	http.HandleFunc("/books/update/process", booksUpdateProcess)
	http.HandleFunc("/books/delete", booksDelete)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/books", http.StatusSeeOther)
}

func booksDelete(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	_, err := db.Exec(`DELETE FROM books WHERE isbn = $1;`, isbn)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/books", http.StatusSeeOther)
}

func booksUpdateProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	bk := Book{}
	bk.ISBN = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("author")
	p := req.FormValue("price") // needs to be float

	if bk.ISBN == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusNotAcceptable)+" Enter valid price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)
	_, err = db.Exec("UPDATE books SET isbn = $1, title=$2, author=$3, price=$4 WHERE isbn= $1;", bk.ISBN, bk.Title, bk.Author, bk.Price)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(res, "updated.gohtml", bk)

}

func booksUpdateForm(res http.ResponseWriter, req *http.Request) {
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
	tpl.ExecuteTemplate(res, "update.gohtml", bk)
}

func booksCreateForm(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "create.gohtml", nil)
}

func booksCreateProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	bk := Book{}
	bk.ISBN = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("author")
	p := req.FormValue("price") // needs to be float

	if bk.ISBN == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusNotAcceptable)+" Enter valid price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)

	// insert
	_, err = db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.ISBN, bk.Title, bk.Author, bk.Price)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(res, "created.gohtml", bk)
}

func booksShow(res http.ResponseWriter, req *http.Request) {
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
	err = tpl.ExecuteTemplate(res, "show.gohtml", bk)
}

func booksIndex(res http.ResponseWriter, req *http.Request) {
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
	err = tpl.ExecuteTemplate(res, "books.gohtml", bks)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
