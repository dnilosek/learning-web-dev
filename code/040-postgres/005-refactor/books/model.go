package books

import (
	"database/sql"
	"errors"

	"github.com/dnilosek/learning-web-dev/code/040-postgres/005-refactor/config"
)

type Book struct {
	ISBN   string  `json:isbn`
	Title  string  `json:title`
	Author string  `json:author`
	Price  float32 `json:price`
}

var ErrRequest = errors.New("Reuest not valid")
var ErrAccept = errors.New("Reuest not acceptable")
var ErrNotFound = errors.New("Not found")
var ErrInternal = errors.New("Internal Error")

func GetOne(isbn string) (Book, error) {
	bk := Book{}
	if isbn == "" {
		return bk, ErrRequest
	}
	row := config.DB.QueryRow(`SELECT * FROM books WHERE isbn = $1;`, isbn)

	err := row.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price) // order matters
	switch {
	case err == sql.ErrNoRows:
		return bk, ErrNotFound
	case err != nil:
		return bk, ErrInternal
	}
	return bk, nil
}

func DeleteOne(isbn string) error {
	if isbn == "" {
		return ErrRequest
	}
	_, err := config.DB.Exec(`DELETE FROM books WHERE isbn = $1`, isbn)

	if err != nil {
		return ErrInternal
	}

	return nil
}

func InsertOne(bk Book) error {
	// insert
	_, err := config.DB.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.ISBN, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return ErrInternal
	}
	return nil
}

func UpdateOne(bk Book) error {
	_, err := config.DB.Exec("UPDATE books SET isbn = $1, title=$2, author=$3, price=$4 WHERE isbn= $1;", bk.ISBN, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return ErrInternal
	}
	return nil
}

func GetAll() ([]Book, error) {
	bks := make([]Book, 0)
	rows, err := config.DB.Query(`SELECT * FROM books;`)
	if err != nil {
		return bks, ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price) // order matters
		if err != nil {
			return bks, ErrInternal
		}
		bks = append(bks, bk)
	}
	return bks, nil
}
