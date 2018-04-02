package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {//{"id":"2","desc":"happy"}

	type Addr struct {
		Id string
		Addesc string
	}
	type Data struct{
		Idq string
		Name string
		//Address Addr
		Address Addr
	}
	addr := Addr{Id:"3",Addesc:"afsdf"}
	content := Data{"6","Kim",addr}
	contentJS,err:=json.Marshal(content)
	if err!=nil{
		panic(err)
	}
	var content2 interface{}
	json.Unmarshal(contentJS,&content2)
	fmt.Println(content2)
	resp, err := http.Post("http://10.0.203.92:8080/heihei", "application/json", bytes.NewReader(contentJS))
	defer resp.Body.Close()
	if err!=nil{
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}