package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

type Data struct {
	Date, Open, High, Low, Close, Volume, adjClose string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln("Cannot read csv")
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV")
	}

	var table []Data
	for _, line := range lines {
		item := Data{
			Date:     line[0],
			Open:     line[1],
			High:     line[2],
			Low:      line[3],
			Close:    line[4],
			Volume:   line[5],
			adjClose: line[6],
		}
		table = append(table, item)
	}

	err = tpl.Execute(os.Stdout, table)
	if err != nil {
		log.Fatalln("Cannot execute template")
	}
}
