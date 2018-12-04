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
	// close the connection
	defer conn.Close()
}

/* Question: golang-web-dev/022_hands-on/02/09_hands-on
Building upon the code from the previous problem:

Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".

Pass the connection of type net.Conn as an argument into this function.

Add "go" in front of the call to "serve" to enable concurrency and multiple connections
*/
