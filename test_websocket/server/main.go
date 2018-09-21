package main

import(
	"golang.org/x/net/websocket"
	"fmt"
	"net/http"
	"flag"
	"log"
)


func init(){
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func handler(conn *websocket.Conn){
	fmt.Printf("a new ws conn: %s->%s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	var err error
	for {
		var reply string
		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			fmt.Println("receive err:",err.Error())
			break
		}
		fmt.Println("Received from client: " + reply)
		if err = websocket.Message.Send(conn, reply); err != nil {
			fmt.Println("send err:", err.Error())
			break
		}
	}
}
func start(addr string)(error){
	http.Handle("/ws", websocket.Handler(ws.handler))
	fmt.Println("begin to listen")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
		return err
	}
	fmt.Println("start end")
	return nil
}

func main(){
	addr  := "127.0.1.1:8888"
	flag.Parse()
	http.Handle("/ws", websocket.Handler(handler))
	fmt.Println("begin to listen")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}
}