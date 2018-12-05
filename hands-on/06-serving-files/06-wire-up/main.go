package main

import (
	"html/template"
	"log"
	"net/http"
)

// set tpl type
var tpl *template.Template

// init template parse glob of everythign in templates/*
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// start main()
func main() {
	// route for index
	http.HandleFunc("/", index)
	// route for about page
	http.HandleFunc("/about", about)
	// route for apply page with button
	http.HandleFunc("/apply", apply)
	// route for contact page
	http.HandleFunc("/contact", contact)

	// start the server
	http.ListenAndServe(":8080", nil)
}

// execute the index template
func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

// execute the about template
func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

// execute the applyProcess or apply template with mux logic based on request method: GET or POST
func apply(w http.ResponseWriter, r *http.Request) {

	// use applyProcess template if request method is POST
	if r.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		if err != nil {
			log.Fatalln("template didn't execute: ", err)
		}
		return
	}

	// use apply template if request method is GET
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

// execute the contact template
func contact(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

/* Question: 024_hands-on/11_hands-on:
Note: skipped 09_hands-on as it and 07 were the same
Starting with the code in the "starting-files" folder:

wire this program up so that it works

ParseGlob in an init function

Use HandleFunc for each of the routes

Combine apply & applyProcess into one func called "apply"

Inside the func "apply", use this code to create the logic to respond differently to a POST method and a GET method
```go
if req.Method == http.MethodPost {
    		// code here
    		return
		}
```

*/
