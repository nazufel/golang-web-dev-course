package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	// http.Handle takes a route and a handler
	// strip prefix from"/resources/" and serve up from "./assets"
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	// should be:
	// http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// serve from "./assets/toby.jpg" since "/resources/" was stripped
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

/*
./assets/toby.jpg
*/
