package main

import (
	"log"
	"os"
	"text/template"
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

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
