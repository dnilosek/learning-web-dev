package books

import (
	"net/http"
	"strconv"

	"github.com/dnilosek/learning-web-dev/code/041-mongodb/config"
)

func BooksDelete(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err := DeleteOne(isbn)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/books", http.StatusSeeOther)
}

func BooksUpdateProcess(res http.ResponseWriter, req *http.Request) {
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
	//bk.Price = float32(f64)
	bk.Price = f64
	err = UpdateOne(bk)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(res, "updated.gohtml", bk)

}

func BooksUpdateForm(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	bk, err := GetOne(isbn)
	switch {
	case err == ErrNotFound:
		http.NotFound(res, req)
		return
	case err != nil:
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(res, "update.gohtml", bk)
}

func BooksCreateForm(res http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(res, "create.gohtml", nil)
}

func BooksCreateProcess(res http.ResponseWriter, req *http.Request) {
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
	//bk.Price = float32(f64)
	bk.Price = f64

	// insert
	err = InsertOne(bk)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(res, "created.gohtml", bk)
}

func BooksShow(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	bk, err := GetOne(isbn)
	switch {
	case err == ErrNotFound:
		http.NotFound(res, req)
		return
	case err != nil:
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = config.TPL.ExecuteTemplate(res, "show.gohtml", bk)
}

func BooksIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	bks, err := GetAll()
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = config.TPL.ExecuteTemplate(res, "books.gohtml", bks)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
