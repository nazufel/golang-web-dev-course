package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		serve(c)
	}
}

func serve(c net.Conn) {
	// close the connection
	defer c.Close()

	// defining variables
	var i int

	// init new scanner to collection connection data
	s := bufio.NewScanner(c)

	// loop over the scanner listening for data
	for s.Scan() {
		// assign data on teh connection to the ln var
		ln := s.Text()
		// print the data to the console
		fmt.Println(ln)

		if i == 0 {
			mux(c, ln)
		}
		if ln == "" {
			// end of the headers
			break
		}
		i++
	}
}

func mux(c net.Conn, ln string) {
	// request line
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	// mux
	if m == "GET" && u == "/" {
		index(c)
	}
	if m == "GET" && u == "/apply" {
		apply(c)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(c)
	}
}

func index(c net.Conn) {
	// code for index function
}

func apply(c net.Conn) {
	// code for apply function with "GET" Request Method
}

func applyProcess(c net.Conn) {
	// code for apply function with "POST" Request Method
}

/* Question: golang-web-dev/022_hands-on/02/17_hands-on
Building upon the code from the previous problem:

Add code to respond to the following METHODS & ROUTES: GET / GET /apply POST /apply
*/
