package main

import "github.com/xtaci/kcp-go"

func main() {
	kcpconn, err := kcp.DialWithOptions("localhost:10000", nil, 10, 3)
	if err!=nil {
		panic(err)
	}

	kcpconn.Write([]byte("hello kcp.emmmmmmmmmmmmmmm"))
	select {}
}
