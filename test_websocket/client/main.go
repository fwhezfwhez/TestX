package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:8888/ws"

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	ws, err := websocket.Dial(url, "", origin)
	//websocket.DialConfig()
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("hello, world!你好")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var buffer = make([]byte, 512)
	var msg = make([]byte,0,512)
	var max int
	var sum int
	for {
		m, err := ws.Read(buffer)
		fmt.Println(max)
		if err != nil {
			if err == io.EOF{
				break
			}
			log.Fatal(err)
		}

		msg = append(msg, buffer[:m]...)
		sum += m
		if sum >= max -4 {
			break
		}
	}
	fmt.Println(len(msg))
	//ws.Close() //关闭连接}

	//var msg2= make([]byte, 512)
	//m2, err2 := ws.Read(msg2)
	//if err2 != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Receive: %s\n", msg2[:m2])
}
