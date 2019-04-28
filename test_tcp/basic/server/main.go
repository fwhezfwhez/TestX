package main

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("tcp run on localhost:7123")
	listener, err := net.Listen("tcp", ":7123")
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	var oneRead []byte
	var e error
	for {
		oneRead, e = readOnce(conn)
		if e != nil {
			if e == io.EOF {
				break
			}
			fmt.Println(errorx.Wrap(e).Error())
			return
		}
		fmt.Println("receive from client:", string(oneRead))

		_, err2 := conn.Write([]byte("reply from server"))

		if err2 != nil {
			fmt.Println(err2.Error())
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readOnce(reader io.Reader) ([]byte, error) {
	var buffer = make([]byte, 512, 512)
	var n int
	var e error

	n, e = reader.Read(buffer)
	if e != nil {

		return nil, e
	}

	return buffer[0:n], nil
}
