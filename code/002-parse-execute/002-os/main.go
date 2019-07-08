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

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Unable to create index.html")
	}

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln("Unable to execute:", err)
	}
}
