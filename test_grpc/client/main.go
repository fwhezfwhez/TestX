package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"test_X/test_grpc/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Username: "ft"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.Message)
}
