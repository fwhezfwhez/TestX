package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	type No struct{
		Url string `json:"url"`
		Data string `json:"data"`
		Interval string `json:"interval"`
	}
	_ , e :=json.Marshal(struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}{Name:"ft", Age:9})
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	buf2,e := json.Marshal(No{
		Url:"http://localhost:8111/get_notify/",
		Data: `{
				"name":"ft"
				}`,
		Interval: "15s",
	})
	c := http.Client{}
	req,e := http.NewRequest("POST","http://localhost:8090/notify/",bytes.NewReader(buf2))
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc4NDgwNzY4MDAsInVzZXJfaWQiOiIxIiwidmVyc2lvbiI6IjIifQ.U0L-h5-XebCaBzGV78_-qfZQFVFIM4JCPcnK6ExTLUI")
	req.Header.Set("Content-Type", "application/json")
	rsp, e :=c.Do(req)
	//defer rsp.Body.Close()
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	b, e:=ioutil.ReadAll(rsp.Body)
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(string(b))
}
