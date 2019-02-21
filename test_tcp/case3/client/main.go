package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var tcpPool chan net.Conn
var receive chan string

func init() {
	receive = make(chan string, 10)
}
func NewClient(connections int, address string) {
	tcpPool = make(chan net.Conn, connections)
	for i := 0; i < connections; i++ {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Panic(err)
		}
		tcpPool <- conn
	}
}

func SendMessage(con net.Conn, msg []byte) error {
	// send message
	_, err := con.Write(msg)
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func ReceiveMessage(con net.Conn) {
	// receiving a message
	inBytes := make([]byte, 0, 1000)
	var b = make([]byte, 512)
	for {
		// bufsize 1024, read bufsize bytes each time
		res, err := con.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
			break
		}
		inBytes = append(inBytes, b[:res]...)
		msg := string(inBytes)
		fmt.Println("receive msg from server:" + msg)
		receive <- msg
	}
}

func getConn() net.Conn {
	con := <-tcpPool
	return con
}

func main() {
	NewClient(20, "localhost:8101")
	con := <-tcpPool
	e := SendMessage(con, []byte("hello, i am client"))
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	go ReceiveMessage(con)
	var msg string
	for {
		select {
		case msg = <-receive:
			fmt.Println(msg)
		}
	}
}
