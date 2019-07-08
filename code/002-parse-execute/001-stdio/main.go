package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("Unable to parse tpl.gohtml:", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Unable to execute:", err)
	}
}
