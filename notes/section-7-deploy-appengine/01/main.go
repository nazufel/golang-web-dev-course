package main

import (
	"net/http"
)

// running on AppEngine does not require func main() or ListenAndServe.
// Use init() to define routes.
func init() {
	http.Handle("/", http.FileServer(http.Dir(".")))
}
