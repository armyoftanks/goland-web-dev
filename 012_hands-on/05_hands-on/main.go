package main

import (
	"log"
	"os"
	"text/template"
)

type items struct {
	Food  string
	Drink string
}

type menu struct {
	Time  string
	Foods []itmes
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	menus := []menu{
		Time: "Breakfast",
		Foods: []items{
			
		}
	}

	err := tpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatalln(err)
	}
}
