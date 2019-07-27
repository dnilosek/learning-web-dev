package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {
	connStr := "postgres://david.nilosek:Nyrod2035@localhost/test?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	defer db.Close()
	check(err)

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/createTable", createTable)
	http.HandleFunc("/dropTable", dropTable)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer (name) VALUES ("Dave")`)
	check(err)
	defer stmt.Closer()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Println(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {

}

func createTable(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func dropTable(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")
}

func index(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT name FROM people;`)
	check(err)

	var s, name string
	s = "RETRIED RECORDS:\n"
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
