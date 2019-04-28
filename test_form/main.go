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
	s :="user_id=fff&password=111"
	req,_:= http.NewRequest("POST","http://localhost:7999/form/", strings.NewReader(s))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rsp,e:=client.Do(req)
	if e!=nil {
		panic(e)
	}
	if rsp.Body!=nil{
		buf,_:=ioutil.ReadAll(rsp.Body)
		fmt.Println(string(buf))
	}

	//rsp,e :=http.PostForm("http://localhost:7999/form/", strings.NewReader(s))
}
