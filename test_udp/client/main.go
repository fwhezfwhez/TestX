package main

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"io"
	"net"
)

var receive = make(chan []byte, 500)

func main() {
	conn, err := net.Dial("udp", "0.0.0.0:7172")
	if err != nil {
		panic(errorx.Wrap(err))
	}
	defer conn.Close()

	go func() {
		for {
			select {
			case tmp := <- Receive(conn):
				fmt.Println("receive from server:", string(tmp))
			}
		}
	}()

	//simple write
	conn.Write([]byte("Hello from client1"))
	conn.Write([]byte("Hello from client2"))
	conn.Write([]byte("Hello from client3"))

	select {}
}

func Receive(conn net.Conn) <-chan []byte {
	var buffer = make([]byte, 500, 500)
	var result = make([]byte, 0, 500)
	for {
		n, e := conn.Read(buffer)
		if e != nil {
			panic(e)
		}
		result = append(result, buffer[0:n]...)
		if e == io.EOF {
			break
		}

		if n < 500 {
			break
		}
	}
	receive <- result
	return receive
}
