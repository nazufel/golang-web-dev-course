package main

import (
	"net/http"
)

func main() {

	// routes
	http.HandleFunc("/", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	// code for dog func
}
