package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// start the server and listen
	li, err := net.Listen("tcp", ":8080")
	// check for errors
	if err != nil {
		log.Fatalln(err)
	}
	// close the connection at the end
	defer li.Close()

	for {
		// accecpt connections
		conn, err := li.Accept()
		// check for errors and print them
		if err != nil {
			log.Println(err)
			continue
		}

		// I like this line at the beginning of the connection, rather than at the end.
		io.WriteString(conn, "I see you connected.\n")

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
				io.WriteString(conn, "\n")
				io.WriteString(conn, "So long!\n")
				io.WriteString(conn, "\n")
				break
			}
		}
		// close the connection
		defer conn.Close()

		// Answer to last question: This string was printed to the console since we were able to break
		// the connection loop with the "logout" string during a telnet session.
		fmt.Println("Code got here.")

		conn.Close()
	}
}

/* Question - golang-web-dev/022_hands-on/02/05_hands-on:

We are now going to get "I see you connected" to be written.

When we used bufio.NewScanner(), our code was reading from an io.Reader that never ended.

We will now break out of the reading.

Package bufio has the Scanner type. The Scanner type "provides a convenient interface for reading data". When you have a Scanner type, you can call the SCAN method on it. Successive calls to the Scan method will step through the tokens (piece of data). The default token is a line. The Scanner type also has a TEXT method. When you call this method, you will be given the text from the current token. Here is how you will use it:

scanner := bufio.NewScanner(conn)
for scanner.Scan() {
	ln := scanner.Text()
	fmt.Println(ln)
}
Use this code to READ from an incoming connection and print the incoming text to standard out (the terminal).

When your "ln" line of text is equal to an empty string, break out of the loop.

Run your code and go to localhost:8080 in your browser.

What do you find?
*/
