package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Animal struct{
	Name string `json:"name"`
	Sex int `json:"sex"`
	Age int `json:"age"`
	Vtype int `json:"type"`
}
type Cat struct{
	Animal
	Color string
}
type Dog struct{
	Animal
	TeethNum int
}

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

	var dog1 = Dog{
		TeethNum: 12,
	}
	dog1.Age = 5
	buf, _ :=json.Marshal(dog1)
	fmt.Println( string(buf))

	t :=reflect.TypeOf(dog1)
	fmt.Println(t.NumField())
}
