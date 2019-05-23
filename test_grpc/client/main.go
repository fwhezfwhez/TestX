package main

import (
	"context"
	"errorX"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"test_X/test_grpc/pb"
	"time"
)

func main() {
	conn, e := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	// say hello
	r, e := c.SayHello(context.Background(), &pb.HelloRequest{Username: "ft"})
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(r.Message)

	// chat
	chatClilent, e :=c.Chat(context.Background())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	go func(){
		for{
			stream, e:=chatClilent.Recv()
			if e == io.EOF {
				fmt.Println("EOF")
				return
			}
			if e != nil {
				fmt.Println(errorx.Wrap(e).Error())
				return
			}
			fmt.Println("receive from server:", stream.Stream)
		}
	}()
	chatClilent.Send(&pb.ClientStream{
		Stream: newBytes(10,9,8,7),
	})
	select{
	 case <-time.After(20 * time.Second):
	}
}

func newBytes(a ...byte)[]byte{
	return a
}
