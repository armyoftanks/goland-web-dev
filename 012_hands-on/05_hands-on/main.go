package main

import (
	"os"
	"text/template"
)

type foods struct {
	Food  []string
	Drink []string
}

type tod struct {
	Time  string
	Foods []foods
}

type menus struct {
	Breakfast tod
	Lunch     tod
	Dinner    tod
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	menu := menus{
		Breakfast: {
			foods{
				Food:  "bacon and eggs",
				Drink: "coffee or OJ",
			},
		},
	}

	err := tpl.Execute(os.Stdout, menu)

}
