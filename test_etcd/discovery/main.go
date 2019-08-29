package main

import (
	"fmt"
	"test_X/test_etcd/discovery/core"
	"time"
)

// 测试前，确保etcd run in 2379
func main() {

	// 模拟服务方将服务注册进etcd
	ser, e := core.NewServiceReg([]string{"localhost:2379"}, 5)
	//ser, e := core.NewServiceReg([]string{"212.64.88.9:2379"}, 5)
	if e!=nil {
		panic(e)
	}
	if e := ser.PutService("/user-login/node/111/", "10.0.1.1:8081"); e != nil {
		panic(e)
	}
	if e := ser.PutService("/user-login/node/112/", "10.0.1.1:8082"); e != nil {
		panic(e)
	}

	time.Sleep(7 * time.Second)

	// 模仿客户端从etcd获取服务路径
	go func(){
		time.Sleep(5 * time.Second)
		cli, e := core.NewClientDis([]string{"localhost:2379"})

		for {
			time.Sleep(5 * time.Second)
			//cli, e := core.NewClientDis([]string{"212.64.88.9:2379"})
			if e!=nil {
				panic(e)
			}
			rs, e := cli.GetService("/user-login/")
			fmt.Println(rs, e)
		}
	}()
	go func(){
		time.Sleep(20 *time.Second)
		ser.DelService("/user-login/node/111/")
	}()
	select {}

}
