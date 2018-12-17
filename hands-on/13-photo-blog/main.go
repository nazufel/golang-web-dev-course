package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
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

	// check and set for image names amd append to cookie var
	c = appendCookie(w,c)

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

func appendCookie(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// set the values for images
	i1 := "disnelyland.jpg"
	i2 := "atbeach.jgp"
	i3 := "hollywood.jpg"

	// append the cookie with a string
	s := c.Value

	// Check if the cookie contains "disneyland.jpg" if not, append
	if !strings.Contains(s, i1) {
		s += "|" + i1
	}
	// Check if the cookie contains "atbeach.jpg" if not, append
	if !strings.Contains(s, i2) {
		s += "|" + i2
	}
	// Check if the cookie contains "hollywood.jpg" if not, append
	if !strings.Contains(s, i3) {
		s += "|" + i3
	}
	// assign appended s to cookie
	c.Value = s
	http.SetCookie(w,c)
	return c
}
