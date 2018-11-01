package main

import (
	"log"
	"os"
	"text/template"
)

type Calihotels []hotel

type hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hotels := Calihotels{
		hotel{
			Name:    "name1",
			Address: "address1",
			City:    "city1",
			Zip:     "11223",
			Region:  "Southern",
		},
		hotel{
			Name:    "name2",
			Address: "address2",
			City:    "city2",
			Zip:     "11223",
			Region:  "Central",
		},
		hotel{
			Name:    "name3",
			Address: "address3",
			City:    "city3",
			Zip:     "11223",
			Region:  "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
