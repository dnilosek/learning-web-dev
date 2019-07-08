package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// Use init to parse files only once
// Use must to fail in init
func init() {
	tpl = template.Must(template.ParseGlob("./*txt"))
}

func main() {

	// Only executes the first file
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Unable to Execute")
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		log.Fatalln("Unable to Execute")
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		log.Fatalln("Unable to Execute")
	}

	// Still executes first one
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Unable to Execute")
	}
}
