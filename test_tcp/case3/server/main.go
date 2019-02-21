package main

import (
	"fmt"
	"io"
	"net"
)

func StartTCPServer(network, addr string) error {
	listener, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {

			fmt.Println(err.Error())
			continue

		}
		onConn(conn)
	}
}

//onConn recieves a tcp connection and waiting for incoming messages
func onConn(conn net.Conn) {
	inBytes := make([]byte, 0)
	// load msg
	for {
		buf := make([]byte, 512)
		res, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err.Error())
			return
		}
		inBytes = append(inBytes, buf[:res]...)

		fmt.Println("receive from client:" + string(inBytes))
		conn.Write([]byte("hello"))
	}
}

func main() {
	if e := StartTCPServer("tcp", ":8101"); e != nil {
		fmt.Println(e.Error())
		return
	}
}
