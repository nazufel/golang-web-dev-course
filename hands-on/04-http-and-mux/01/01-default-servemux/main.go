package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Index")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dogo!")
}

func me(res http.ResponseWriter, req *http.Request) {
	m := "Hello Ryan!"
	io.WriteString(res, m)
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

/* Question: golang-web-dev/022_hands-on/01/01_hands-on

ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/
