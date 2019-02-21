package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	user "test_X/go-micro-user-validate/service/example_user/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.api.user"))
	service.Init()

	// Create new greeter client
	userService := user.NewUserService("go.micro.api.user", service.Client())

	// Call the greeter
	rsp, err := userService.GetToken(context.TODO(), &user.Request{Username: "John", Password: "123"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp)
}
