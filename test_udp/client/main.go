package main

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("udp", "0.0.0.0:7172")
	if err != nil {
		panic(errorx.Wrap(err))
	}
	defer conn.Close()

	go Receive(conn)

	for i :=0;i<30;i++ {
		conn.Write([]byte("Hello,I am server" + strconv.Itoa(i)))

	}
	//conn.Write([]byte("Hello from client1"))

	select {}
}

func Receive(conn net.Conn) {
	var buffer = make([]byte, 500, 500)

	for {

		n, e := conn.Read(buffer)
		if e != nil {
			panic(e)
		}
		fmt.Println("receive from server: ", string(buffer[0:n]))
	}
}
