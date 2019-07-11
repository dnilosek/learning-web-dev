package main

import (
	"log"
	"os"
	"text/template"
)

type Menu struct {
	Type      string
	FoodItems []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	menus := []Menu{
		Menu{
			Type: "Breakfast",
			FoodItems: []string{
				"Eggs",
				"Bacon",
				"NotEggs",
			},
		},
		Menu{
			Type: "Lunch",
			FoodItems: []string{
				"Sammich",
				"More Sammich",
			},
		},
		Menu{
			Type: "Dinner",
			FoodItems: []string{
				"Steak",
				"More Steak",
			},
		},
	}

	err := tpl.Execute(os.Stdout, menus)
	if err != nil {
		log.Fatalln("Unable to execute")
	}
}
