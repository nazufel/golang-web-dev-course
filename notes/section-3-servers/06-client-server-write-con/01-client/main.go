package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
	/*
		bs, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(bs))
	*/
}
