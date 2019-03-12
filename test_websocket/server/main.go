package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func handler(conn *websocket.Conn) {
	fmt.Printf("a new ws conn: %s->%s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	var err error
	for {
		var receive []byte
		// 不需要为receive初始化，因为Receive内置了
		err = websocket.Message.Receive(conn, &receive)
		fmt.Println(receive)
		if err != nil {
			if err == io.EOF {
				fmt.Println("结束")
				break
			}else{
				fmt.Println("receive err:", err.Error())
				break
			}
		}
		fmt.Println("Received from client: " + string(receive))

		var reply = make([]byte,2049)
		reply[0] = 1
		reply[2048] = 1
		fmt.Println(len(reply))
		if err = websocket.Message.Send(conn, reply); err != nil {
			fmt.Println("send err:", err.Error())
			break
		}
	}
}

func main() {
	addr := "127.0.0.1:8888"
	flag.Parse()
	http.Handle("/ws", websocket.Handler(handler))
	fmt.Println("begin to listen")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}
}
