package main

import (
	"html/template"
	"log"
	"net/http"
)

// I cheated on this one, but I learned so much and this is probabaly be the design pattern moving forward.
// This is what I was trying to do earlier.

// Step 2: set tpl type
var tpl *template.Template

// Step 3: init the index templates
func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

// Step 1: start main()
func main() {
	// Step 4: requests to the "/" route runs the "dogs" func
	http.HandleFunc("/", dogs)
	// Step 6: index.gohtml templates made calls to "/resources/", which then served up all files in the public dir.
	// the root is "/resources/" in the project but StripPrefix removes that and services everything from "public"
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	// Step 7: start the server
	http.ListenAndServe(":8080", nil)
}

// Step 5: dogs func executes the templates, which then makes calls to "/resources/"
func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("templates didn't execute: ", err)
	}
}

/* Question: 024_hands-on/07_hands-on:

Serve the files in the "starting-files" folder

To get your images to serve, use:

	func StripPrefix(prefix string, h Handler) Handler
	func FileServer(root FileSystem) Handler
Constraint: you are not allowed to change the route being used for images in the templates file

*/
