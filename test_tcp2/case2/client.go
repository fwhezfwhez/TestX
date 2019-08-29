package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"test_X/test_tcp/case2/pb"
	"zonst/qipai/protocolutil"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func main() {
	hostInfo := "0.0.0.0:8113"
	conn, err := net.Dial("tcp", hostInfo)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		log.Println(err.Error())
		return
	}
	helloRequest := pb.HelloRequest{
		MessageId: int32(999),
		Name:      "fengtao",
	}
	header := &protocolutil.ClientHeader{
		MessageType: protocolutil.ClientRequestMessageType,
		DestID:      0,
		MessageID:   uint16(helloRequest.MessageId),
	}

	packet := protocolutil.NewPacket(header)
	packet.SetContent(&helloRequest)
	buf := packet.Serialize()
	conn.Write(buf)
	rs, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("rs:", rs)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
