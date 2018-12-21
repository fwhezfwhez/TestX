package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
type Form struct{
	ViewState string `form`
}
var client = http.Client{}
func main() {
	s :="username=fff&password=111"
	req,_:= http.NewRequest("POST","http://www.baidu.com", strings.NewReader(s))
	rsp,e:=client.Do(req)
	if e!=nil {
		panic(e)
	}
	if rsp.Body!=nil{
		buf,_:=ioutil.ReadAll(rsp.Body)
		fmt.Println(string(buf))
	}
}
