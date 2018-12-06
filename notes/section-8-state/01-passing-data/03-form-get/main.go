package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// form value to be returned in the URL: /?n=<requested value>
	v := req.FormValue("n")
	// set header for response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// getting information through the URL: localhost:8080/?q=
	io.WriteString(w, `
	<form method="get">
	 <input type="text" name="n">
	 <input type="submit">
	</form>
	<br>`+v)
}
