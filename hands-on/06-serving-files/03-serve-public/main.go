package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	// use the http.FileServer for anything after "/pics/"
	http.Handle("/pic/", fs)
	http.ListenAndServe(":8080", nil)
}

/* Question: 024_hands-on/05_hands-on:

Serve the files in the "starting-files" folder

To get your images to serve, use only this:

	fs := http.FileServer(http.Dir("public"))
Hint: look to see what type FileServer returns, then think it through.
*/
