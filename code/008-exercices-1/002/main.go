package main

import (
	"log"
	"os"
	"text/template"
)

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type HotelRegion struct {
	Region string
	Hotels []Hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	n1 := Hotel{}
	n2 := Hotel{}
	northernRegion := HotelRegion{
		Region: "Northern",
		Hotels: []Hotel{n1, n2},
	}

	c1 := Hotel{}
	c2 := Hotel{}
	centralRegion := HotelRegion{
		Region: "Central",
		Hotels: []Hotel{c1, c2},
	}

	s1 := Hotel{}
	s2 := Hotel{}
	sourthernRegion := HotelRegion{
		Region: "Southern",
		Hotels: []Hotel{s1, s2},
	}

	data := []HotelRegion{
		sourthernRegion,
		centralRegion,
		northernRegion,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln("Unable to Execute template")
	}
}
