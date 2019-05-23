package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"test_X/test_grpc/pb"
)

type HelloService struct {
}

func (hs HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("你好，%s", in.Username)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":6001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, HelloService{})
	s.Serve(lis)
}
