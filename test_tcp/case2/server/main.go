package main

import (
	"net"
	"os"
	"fmt"
)

func main() {

	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// run as a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()

	var buf [512]byte
	for {
		// read upto 512 bytes
		_, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// write the n bytes read
		_, err2 := conn.Write([]byte("reply from server"))
		if err2 != nil {
			fmt.Println(err2.Error())
			continue
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}