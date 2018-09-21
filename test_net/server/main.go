package main

import (
	"net"
	"fmt"
)


func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		conn.
		go handleConnection(conn)
	}
}
