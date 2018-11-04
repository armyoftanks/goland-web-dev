package main

import (
	"log"
	"os"
	"text/template"
)

type foods struct {
	Food  string
	Drink string
}

type tod struct {
	Time  string
	Foods []foods
}

type menus struct {
	Tod tod
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	menu := menus{
		tod{
			Time: "Breakfast: ",
			foods{
				Food:  "bacon and eggs",
				Drink: "coffee or OJ",
			},
		},
	}

	err := tpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatalln(err)
	}
}
