package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.txt")
	if err != nil {
		log.Fatalln("Unable to parse one.txt")
	}
	//tpl.Execute(os.Stdout, nil)

	tpl, err = tpl.ParseFiles("two.txt", "three.txt")
	if err != nil {
		log.Fatalln("Unable to parse remaining files")
	}

	// Executes in LIFO order
	err = tpl.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		log.Fatalln("Unable to execute template")
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		log.Fatalln("Unable to execute template")
	}

	tpl.Execute(os.Stdout, nil)
}
