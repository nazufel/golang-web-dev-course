package main

import (
	"fmt"
	"net/http"
)

func main() {

	// route for root
	http.HandleFunc("/", index)
	// handling favicon requests since I don't have one
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// print the path requested in the terminal
	fmt.Println(r.URL.Path)
	// write back to the connection to check the temrminal
	fmt.Fprintln(w, "go look at your terminal")
}
