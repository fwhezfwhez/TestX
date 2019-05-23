package main

import (
	"fmt"
	"github.com/xtaci/kcp-go"
)

func main() {
	kcpconn, err := kcp.DialWithOptions("localhost:10000", nil, 10, 3)
	if err!=nil {
		panic(err)
	}

	go func ()(){
		var buffer = make([]byte,1024,1024)
		for {
			n ,e :=kcpconn.Read(buffer)
			if e!=nil {
				fmt.Println(e.Error())
				break
			}
			fmt.Println(string(buffer[:n]))
		}

	}()
	kcpconn.Write(NewBytes(1,2,3,4,5,6))

	//time.Sleep(5 * time.Second)
	kcpconn.Write(NewBytes(7,8,9,10,11))

	select {}
}

func NewBytes(a ... byte) []byte {
	return a
}
