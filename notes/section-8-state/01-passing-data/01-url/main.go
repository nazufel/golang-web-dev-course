package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// putting value into URL. This would be done by going to
	// localhost:8080/?q=<something to print>
	v := req.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)
}

// visit this page:
// http://localhost:8080/?q=dog
