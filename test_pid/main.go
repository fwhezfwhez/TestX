package main

import (

	"fmt"

	"errors"
)

func main() {

	//m :=make(map[string]string)
	//fmt.Println(m["1"]=="")
	//if pid := syscall.Getpid(); pid != 1 {
	//	//ioutil.WriteFile("server.pid", []byte(strconv.Itoa(pid)), 0777)
	//	//defer os.Remove("server.pid")
	//	fmt.Println(strconv.Itoa(pid))
	//}
	//
	//// 捕获kill的信号
	//sigTERM := make(chan os.Signal, 1)
	//signal.Notify(sigTERM, syscall.SIGTERM)
	//// 收到信号前会一直阻塞
	//<-sigTERM
	//fmt.Print("killed")



	fmt.Println(fmt.Sprintf("%v,%v,%v",errors.New("错误1"),errors.New("错误2"),errors.New("错误3")))
}