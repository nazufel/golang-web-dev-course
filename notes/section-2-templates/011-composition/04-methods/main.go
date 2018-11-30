package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

func (p person) Process() int {
	return 9
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) Args(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p := person{
		Name: "James Bond",
		Age:  42,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
