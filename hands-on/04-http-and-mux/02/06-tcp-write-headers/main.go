package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	// close the connection
	defer conn.Close()

	// I like this line at the beginning of the connection, rather than at the end.
	io.WriteString(conn, "\nI see you connected.\n")
	io.WriteString(conn, "\nType something to the TCP connection.\n")

	// init new scanner to collection connection data
	scanner := bufio.NewScanner(conn)
	// scan for data on the connection
	for scanner.Scan() {
		// assign data on the connection to the lh var
		ln := scanner.Text()
		// print out the data coming in
		fmt.Println(ln)

		// close the loop when ln data == "logout"
		if ln == "logout" {
			// This code from the last hands-on fulfills this one as well
			// since it writes back to the connection with, "\n So long!\n\n"
			io.WriteString(conn, "\nSo long!\n\n")
			break
		}
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

/* Question: golang-web-dev/022_hands-on/02/11_hands-on
Building upon the code from the previous problem:

Before we WRITE our RESPONSE , let's WRITE to our RESPONSE the STATUS LINE and some REPONSE HEADERS. Remember the request line and status line:

REQUEST LINE GET / HTTP/1.1 method SP request-target SP HTTP-version CRLF https://tools.ietf.org/html/rfc7230#section-3.1.1

RESPONSE (STATUS) LINE HTTP/1.1 302 Found HTTP-version SP status-code SP reason-phrase CRLF https://tools.ietf.org/html/rfc7230#section-3.1.2

Write the following strings to the response - use io.WriteString for all of the following except the second and third:

"HTTP/1.1 200 OK\r\n"

fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))

fmt.Fprint(c, "Content-Type: text/plain\r\n")

"\r\n"

Look in your browser "developer tools" under the network tab. Compare the RESPONSE HEADERS from the previous file with the RESPONSE HEADERS in your new solution.
*/
