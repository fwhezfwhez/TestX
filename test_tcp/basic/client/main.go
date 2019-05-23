package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var receive = make(chan []byte, 512)

func main() {
	hostInfo := "127.0.0.1:7123"
	conn, err := net.Dial("tcp", hostInfo)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		log.Println(err.Error())
		return
	}
	go func() {
		for {
			recv, e := Receive(conn)
			if e != nil {
				if e == io.EOF {
					break
				}
				panic(e)
			}
			select {
			case rs := <-recv:
				fmt.Println(fmt.Sprintf("receive message from server side: %v", rs))
			}
		}
	}()

	conn.Write([]byte(NewByte(1,2,3,4,5,6)))
	conn.Write([]byte(NewByte(1,2,3,4,5,6)))
	conn.Write([]byte(NewByte(1,2,3,4,5,6)))

	select {}
}

func Receive(conn net.Conn) (<-chan []byte, error) {
	var buffer = make([]byte, 512, 512)
	var n int
	var e error
	n, e = conn.Read(buffer)
	if e != nil {
		return nil, e
	}

	receive <- buffer[0:n]
	return receive, nil
}

func NewByte(byts ...byte) []byte {
	var rs = make([]byte, 0 , 512)
	rs = append(rs, byts...)
	return rs
}
