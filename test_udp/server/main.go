package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":7172")
	fmt.Println("start udp on 7172")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// listen to incoming udp packets
	go func(packetConn net.PacketConn) {
		var times int
		for {
			// read to EOF
			message, addr := ReadToEOF(conn)
			times ++
			fmt.Println("times:", times)
			fmt.Println("receive from client:", string(message))
			//simple write
			conn.WriteTo([]byte("Hello,I am server"), addr)
			conn.WriteTo([]byte("Hello,I am server2"), addr)
			continue
		}
	}(conn)

	select {}
}

func ReadToEOF(conn net.PacketConn) ([]byte, net.Addr) {
	var buffer = make([]byte, 4096, 4096)

	n, addr, e := conn.ReadFrom(buffer)
	fmt.Println(n)

	if e != nil {
		panic(e)
	}
	return buffer[0:n], addr
}
