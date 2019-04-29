package main

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"net"
)

var receive = make(chan []byte, 500)

func main() {
	conn, err := net.Dial("udp", "0.0.0.0:7172")
	if err != nil {
		panic(errorx.Wrap(err))
	}
	defer conn.Close()

	go Receive(conn)
	go func() {
		for {
			select {
			case tmp := <-receive:
				fmt.Println("receive from server:", string(tmp))
			}
		}
	}()

	//simple write
	conn.Write([]byte("Hello from client1"))

	select {}
}

func Receive(conn net.Conn) {
	var buffer = make([]byte, 500, 500)
	for {
		n, e := conn.Read(buffer)
		if e != nil {
			panic(e)
		}
		receive <- buffer[0:n]
	}
}
