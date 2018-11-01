package main

import (
  "text/template"
  "os"
  "log"
)

var tpl *template.Template

type Calihotels struct {
  Hotels []hotel
}

type hotel struct {
  Name string
  Adress string
  City string
  Zip int
  Region string
}

func init {
  tpl = template.MUST(template.ParseFiles("tpl.gohtml"))
}

func main {

  hotels := []Calihotels {
    hotel{"name1", "address1", "city1", 11223, "Southern"},
    hotel{"name2", "address2", "city2", 11223, "Central"},
    hotel{"name3", "address3", "city3", 11223, "Northern"},
  }

  

}
