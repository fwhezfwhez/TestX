package main

import (
	"google.golang.org/grpc"

	pb "micro-rbac/proto"

	"golang.org/x/net/context"

	"fmt"

)

func main() {
	//client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	//if err != nil {
	//	fmt.Println("链接rpc服务器失败:", err)
	//}
	//var reply int
	//err = client.Call("Watcher.GetInfo", 9, &reply)
	//if err != nil {
	//	fmt.Println("调用远程服务失败", err)
	//}
	//fmt.Println("远程服务返回结果：", reply)


	rbac, _ := grpc.Dial("localhost:13030", grpc.WithInsecure())

	client := pb.NewRbacServiceClient(rbac)

	// 登录 ok

	//reply, er:= client.Login(context.Background(), &pb.LoginReq{Username: "752825581@qq.com",Password:"123456",PlatformId:10000})
	//if er!=nil{
	//	fmt.Println(er.Error())
	//	return
	//}
	//
	//
	//if reply.Error.Code != 0 {
	//
	//	fmt.Println("login",reply.Error.Code, reply.Error.Message)
	//
	//}
	//fmt.Println(reply)
	// 注销 ok
	//
	//reply2,_:=client.Logout(context.Background(),&pb.LogoutReq{Userid:10062,PlatformId:10000})
	//
	//if reply2.Error.Code != 0 {
	//
	//	fmt.Println("logout",reply2.Error.Code, reply2.Error.Message)
	//
	//}
	//fmt.Println(reply2)
	//后台actionplateformid-4，前台也是10000

	//注册
	//reply2,_:=client.SignUp(context.Background(),&pb.SignUpReq{PlatformId:10000,Username:"14379672@qq.com",Password:"123456",ActionPlatformId:4,ActionUserId:10039})
	//
	//if reply2.Error.Code != 0 {
	//
	//	fmt.Println("logout",reply2.Error.Code, reply2.Error.Message)
	//
	//}
	//fmt.Println(reply2)
	// 获取用户

	//reply3,_:=client.GetUsers(context.Background(),&pb.UserReq{Username:"752825581@qq.com",PlatformId:10000})
	//
	//if reply3.Error.Code != 0 {
	//
	//	fmt.Println("users",reply3.Error.Code, reply3.Error.Message)
	//
	//}else {
	//
	//	fmt.Println(reply3.Users)
	//
	//}

	// 注册
	//undefine
	//验证token

//	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjgzNDE3MDYsImlkIjoxMDA2MiwidXNlcm5hbWUiOiI3NTI4MjU1ODFAcXEuY29tIiwidmVyc2lvbiI6ImU4MTExNjQxLTk1OTUtNDY5Yy04NGQxLTcyNzVlNTFhOGEzMCJ9.jCbs49UKMA37IRNGwr8A4KC_O42DDPxfZCwDcGnZZ-U
	reply4,_:=client.CheckToken(context.Background(),&pb.TokenReq{Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjgzNzQ1ODcsImlkIjoxMDA2MiwidXNlcm5hbWUiOiI3NTI4MjU1ODFAcXEuY29tIiwidmVyc2lvbiI6ImU4MTExNjQxLTk1OTUtNDY5Yy04NGQxLTcyNzVlNTFhOGEzMCJ9.7dIjlxxj6S-AV4Kr9rj9-7S6iVEvgPRtub0RYEmXQo8",PlatformId:10000})

	if reply4.Error.Code != 0 {

		fmt.Println("checktoken",reply4.Error.Code, reply4.Error.Message)

	}

	fmt.Println(reply4)
	//密码修改

	//reply5,_:=client.PassWordUpdate(context.Background(),&pb.PassWordReq{Username:"1020300659@qq.com",
	//
	//	Password:"123456",ActionUserId:10041,ActionPlatformId:4})
	//
	//if reply5.Error.Code != 0 {
	//
	//	fmt.Println("password",reply5.Error.Code, reply5.Error.Message)
	//
	//}


f(f3)
}

func f3(int){}
func f(f2 func(interface{})){

}
