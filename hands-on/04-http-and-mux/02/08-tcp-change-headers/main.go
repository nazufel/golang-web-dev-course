package main

import (
	"bufio"
	"fmt"
	"io"
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
	var rMethod, rURI string

	// init new scanner to collection connection data
	s := bufio.NewScanner(c)

	// loop over the scanner listening for data
	for s.Scan() {
		// assign data on teh connection to the ln var
		ln := s.Text()
		// print the data to the console
		fmt.Println(ln)

		if i == 0 {
			// split the data string fields at the point of white space and assign to xs as a slice of strings
			xs := strings.Fields(ln)
			// first string field is the request method
			rMethod = xs[0]
			// second string field is the requst URI
			rURI = xs[1]
			// print the request method
			fmt.Println("METHOD:", rMethod)
			// print the URI method
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("End headers.")
			break
		}
		i++
	}

	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				<strong>HOLY COW THIS IS LOW LEVEL</strong>
			</body>
			</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

/* Question: golang-web-dev/022_hands-on/02/15_hands-on
Building upon the code from the previous problem:

Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in

tags.
*/
