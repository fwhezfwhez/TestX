package main

import (
	"bufio"
	"fmt"
	"golang.org/x/protobuf/proto"
	"log"
	"net"
	"test_X/test_tcp/case2/pb"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func main() {
	hostInfo := "0.0.0.0:8113"
	conn, err := net.Dial("tcp", hostInfo)
	if conn !=nil {
		defer conn.Close()
	}
	if err != nil {
		log.Println(err.Error())
		return
	}
	helloRequest:= pb.HelloRequest{
		MessageId: int32(999),
		Name: "fengtao",
	}

	buf ,e:=proto.Marshal(&helloRequest)
	if e!=nil {
		log.Println(e.Error())
		return
	}
	conn.Write([]byte(buf))
	rs, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("rs:", rs)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
