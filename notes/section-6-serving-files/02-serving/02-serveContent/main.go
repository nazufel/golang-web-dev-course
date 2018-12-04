package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// look at the absolute path for the picture
	io.WriteString(w, `
	<img src="/toby.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {

	// open the file.
	f, err := os.Open("toby.jpg")

	// throw 404 is the file isn't there
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	// defer closing
	defer f.Close()

	// serve the file by copying it to the response writer
	io.Copy(w, f)
}
