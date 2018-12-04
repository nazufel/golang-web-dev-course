package main

import (
	"io"
	"net/http"
)

func main() {

	// routes
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// path to the file
	io.WriteString(w, `<img src="toby.jpg">	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {

	// take the request, serve the file to the response
	http.ServeFile(w, req, "toby.jpg")
}
