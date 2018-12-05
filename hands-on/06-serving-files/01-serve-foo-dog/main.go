package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	//http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

/* Todd's question said nothing about having another func. I was not able to
	execute the tempalte and serve the file with http.ServeFile. However, I disagree
	with using http.ServeFile in a separate func. Let the template handle that as I have done.
	Leaving this commented out since I may need it later as the hands-on build on eachother.

func chien(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

/* Question: 024_hands-on/03_hands-on:

Serve the files in the "starting-files" folder

Use "http.FileServer"
*/
