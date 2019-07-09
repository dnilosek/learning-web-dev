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

type car struct {
	Make  string
	Model string
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
	s2 := sage{
		Name:  "God",
		Motto: "I said so",
	}
	sages := []sage{s1, s2}

	c1 := car{
		Make:  "Ford",
		Model: "F150",
	}
	c2 := car{
		Make:  "Mazda",
		Model: "Mazda3",
	}
	cars := []car{c1, c2}

	data := struct {
		Cars  []car
		Sages []sage
	}{cars, sages}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln("Unable to execute template")
	}
}
