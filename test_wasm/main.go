package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 1.打印 hello wasm
	fmt.Println("hello,wasm")

	// 发送 http请求
	// httpClient()
	fmt.Println(CalculateAdd(1,2,3,4,5))
}

func CalculateAdd(a...int) int{
	sum :=0
	for _,v:=range a{
		sum+=v
	}
	return sum
}

func httpClient(){
	c:= http.Client{}
	req,e:= http.NewRequest("GET","http://localhost:8889/hello/",nil)
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	rsp,e:=c.Do(req)
	if rsp.Body!=nil{
		defer rsp.Body.Close()
	}
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	buf,e := ioutil.ReadAll(rsp.Body)
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(string(buf))
}