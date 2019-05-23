package main

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"github.com/xtaci/kcp-go"
	"io"
	"net"
)

func main() {
	fmt.Println("kcp listens on 10000")
	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)
	if err!=nil {
		panic(err)
	}
	for {
		conn, e :=lis.AcceptKCP()
		if e!=nil {
			panic(e)
		}
		go func(conn net.Conn){
			var buffer = make([]byte,1024,1024)
			for {
				n,e :=conn.Read(buffer)
				if e!=nil {
					if e == io.EOF {
						break
					}
					fmt.Println(errorx.Wrap(e))
					break
				}

				// 开启协程处理客户端请求，防止一条请求未处理完时，另一条请求阻塞
				go func () {
					fmt.Println("receive from client:", buffer[:n])
					conn.Write([]byte("你好"))
				}()
			}
		}(conn)
	}


}
