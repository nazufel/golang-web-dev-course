package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	// get the cookie by running the getCookie func
	c := getGooke(w, r)


	// process the submission form
	if r.Method == http.MethodPost {
		// get multipart file, file header and err
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		// close the file when done
		defer mf.Close()

		// pull out the extention value since strings.Split, splits a stirng
		// into a slice of strings. the [1] is the "second" value split on
		// on the ".".
		ext := strings.Split(fh.Filename, ".")[1]

		// create sha for file
		h := sha1.New()

		// copy the file name to the hashed file name
		io.Copy(h, mf)

		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// get the working directory
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		// join the path for creating a file
		p := filepath.Join(wd, "public", "pics", fname)

		// create the file
		nf, err := os.Create(p)
		if err != nil {
			fmt.Println(err)
		}
		// close the new file
		defer nf.Close()

		// copy contents of old file to new
		mf.Seek(0, 0)
		io.Copy(nf, mf)

		// add the file name to the user's cookie
		c = appendCookie(w, c, fname)
	}
	// split out the strings into a slice of strings
	xs := strings.Split(c.Value, "|")

	// execute the templates after getting the cookie
	tpl.ExecuteTemplate(w, "index.gohtml", xs)

}

// getCookie func checks if the client has a cookie, if not then make one. Lastly, return the cookie.
func getGooke(w http.ResponseWriter, r *http.Request) *http.Cookie {

	// get the session cookie from the request
	c, err := r.Cookie("session")
	// check the error and if there is no cookie, make one
	if err != nil {
		// generate new cookie string using uuid
		sID, _ := uuid.NewV4()
		// create a cookie in the client with the name "session" and the uuid value
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		// write the cookie to the client
		http.SetCookie(w, c)
	}
	// return the cookie value out of the func and back into the calling index func
	return c
}

func appendCookie(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	// append the cookie with a string
	s := c.Value

	// Check if the cookie contains the image string if not, append
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}

	// assign appended s to cookie
	c.Value = s
	http.SetCookie(w,c)
	return c
}
