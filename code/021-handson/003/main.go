package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var tpl *template.Template

func monthDayYear(t time.Time) string {
	return t.Format(time.RFC850)
}

var funcMap = template.FuncMap{
	"formatDate": monthDayYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("tpl.gohtml"))
}

func root(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(root))

	http.ListenAndServe(":8080", nil)
}
