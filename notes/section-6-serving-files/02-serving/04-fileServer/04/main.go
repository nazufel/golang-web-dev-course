package main

import (
	"log"
	"net/http"
)

// having index.html rather than index.gohtml will render the static site rather than serve the files
func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
