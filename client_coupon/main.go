package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

var host = "http://localhost:8077"
var client = http.Client{}
func main() {
	req,er := http.NewRequest("GET",host+"/xmCoupon",nil)
	if er!=nil{
		fmt.Println(er)
		return
	}
	resp,er:=client.Do(req)
	if er!=nil{
		fmt.Println(er)
		return
	}
	helpRead(resp)
}


func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}