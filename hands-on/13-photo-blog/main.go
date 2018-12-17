package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
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

	// execute the templates after getting the cookie
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)

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
	// return the cookie value out of the func
	return c
}
