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

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

/* Question: golang-web-dev/022_hands-on/01/05_hands-on

1. Take the previous program and change it so that:
* func main uses http.Handle instead of http.HandleFunc

Contstraint: Do not change anything outside of func main

Hints:

http.HandlerFunc

type HandlerFunc func(ResponseWriter, *Request)
http.HandleFunc

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
source code for HandleFunc

  func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
  		mux.Handle(pattern, HandlerFunc(handler))
  }
*/
