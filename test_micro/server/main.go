package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	hello_world "test_X/test_micro/model"
)

type HelloWorld struct{}

func (g *HelloWorld) Hello(ctx context.Context, req *hello_world.HelloRequest, rsp *hello_world.HelloResponse) error {
	rsp.Greeting = "Hello World: " + req.Name
	return nil
} // 实现hello_world service中Hello方法

func main() {
	service := micro.NewService(
		micro.Name("hello_world"), // 定义service的名称为hello_world
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init() // 初始化service
	micro.RegisterHandler(service.Server(), new(HelloWorld)) // 注册服务
	fmt.Println(service.Options())
	if err := service.Run(); err != nil {
		fmt.Println(err)
	} // 运行服务

}
