package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:8888/ws"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	websocket.DialConfig()
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("hello, world!你好")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	message2 := []byte("第二条msg")
	_, err = ws.Write(message2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message2)

	var msg = make([]byte, 512)
	m, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])
	//ws.Close() //关闭连接}
	select {}

	//var msg2= make([]byte, 512)
	//m2, err2 := ws.Read(msg2)
	//if err2 != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Receive: %s\n", msg2[:m2])
}
