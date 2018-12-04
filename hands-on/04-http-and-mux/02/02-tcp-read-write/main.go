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

		// init new scanner to collection connection data
		scanner := bufio.NewScanner(conn)
		// scan for data on the connection
		for scanner.Scan() {
			// assign data on the connection to the lh var
			ln := scanner.Text()
			// print out the data coming in
			fmt.Println(ln)
		}
		// close the connection
		defer conn.Close()

		// Answer to last question: the code never gets here becuase the for loop is never closed.
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}

/* Question - golang-web-dev/022_hands-on/02/03_hands-on:

Building upon the code from the previous exercise:

In that previous exercise, we WROTE to the connection.

Now I want you to READ from the connection.

You can READ and WRITE to a net.Conn as a connection implements both the reader and writer interface.

Use bufio.NewScanner() to read from the connection.

After all of the reading, include these lines of code:

fmt.Println("Code got here.") io.WriteString(c, "I see you connected.")

Launch your TCP server.

In your web browser, visit localhost:8080.

Now go back and look at your terminal.

Can you answer the question as to why "I see you connected." is never written?
*/
