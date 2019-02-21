package main

import (
	"context"
	"github.com/micro/go-api"
	"github.com/micro/go-api/handler/http"
	"github.com/micro/go-micro"
	"test_X/go-micro-user-validate/service/example_user/proto"
	"time"
)

type User struct {
}

func (u *User) GetToken(c context.Context, req *user.Request, rsp *user.Response) error {
	rsp.Exp = time.Now().Add(24 * time.Hour).Local().String()
	rsp.Token = "Bearer ea1jf2ad253d1da8a7d0f9g98ad9z9fas09dkv5.24csad.142.sc"
	rsp.HttpResponse = &user.HttpResponse{}
	rsp.HttpResponse.StatusCode=400
	rsp.HttpResponse.Body = "hello"
	return nil
}

func (u *User) GetUserName(ctx context.Context, req *user.Request, rsp *user.UserName) error {
	rsp.Username = "Ft"

	rsp.Message = "fail"
	//return errors.BadRequest("go.micro.api.user", "hehe")
	return nil
}
func main() {
	userService := micro.NewService(micro.Name("go.micro.api.user"))
	userService.Init()

	//// or
	//user.RegisterUserHandler(userService.Server(), new(User),
	//	api.POST("/user/token/", "User.GetToken"),
	//	api.GET("/user/username/", "User.GetUserName"),
	//)
	//
	// or
	 user.RegisterUserHandler(userService.Server(), new(User), api.WithEndpoint(&api.Endpoint{
		  Name: "User.GetToken",
		  Path: []string{"/user/token/"},
		  Method: []string{"POST"},
		  Handler: http.Handler,
	 }),api.WithEndpoint(&api.Endpoint{
		  Name: "User.GetUserName",
		  Path: []string{"/user/username/"},
		  Method: []string{"GET"},
		  Handler: http.Handler,
	 }))

	//
	// or
	//r := api.NewHttpRouters()
	//r.POST("/user/token/", "User.GetToken")
	//r.GET("/user/username/", "User.GetUserName")
	//
	//user.RegisterUserHandler(userService.Server(), new(User), r...)

	if er := userService.Run(); er != nil {
		panic(er)
	}
}
