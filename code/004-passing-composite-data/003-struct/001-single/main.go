package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	s1 := sage{
		Name:  "James Bond",
		Motto: "Shaken, not stirred",
	}

	err := tpl.Execute(os.Stdout, s1)
	if err != nil {
		log.Fatalln("Unable to execute template")
	}
}
