package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func doTask(conn net.Conn) {
	//
	//conn.
}
func main() {
	hostInfo := "127.0.0.1:1201"
	conn, err := net.Dial("tcp", hostInfo)
	if conn !=nil {
		defer conn.Close()
	}
	if err != nil {
		log.Println(err.Error())
		return
	}
	conn.Write([]byte("request from client"))
	rs, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("rs:", rs)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//conn.Write([]byte("helloWorld2"))
	//rs, err = bufio.NewReader(conn).ReadString('\n')
	//fmt.Println(rs)
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//for{
	//	conn.Write([]byte("helloWorld"))
	//	rs, err := bufio.NewReader(conn).ReadString('\n')
	//	if err != nil {
	//		log.Println(err.Error())
	//		return
	//	}
	//	fmt.Println(rs)
	//}

}
