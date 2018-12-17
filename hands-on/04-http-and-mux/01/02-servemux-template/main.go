package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func index(res http.ResponseWriter, req *http.Request) {

	d := "index"

	err := tpl.Execute(res, d)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(res http.ResponseWriter, req *http.Request) {

	d := "Dogo!"

	err := tpl.Execute(res, d)
	if err != nil {
		log.Fatalln(err)
	}
}

func me(res http.ResponseWriter, req *http.Request) {

	d := "Me"

	err := tpl.Execute(res, d)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

/* Question: golang-web-dev/022_hands-on/01/03_hands-on

1. Take the previous program in the previous folder and change it so that:

* a templates is parsed and served

* you pass data into the templates
*/
