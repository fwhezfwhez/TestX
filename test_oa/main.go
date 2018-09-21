package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){
	req,er :=http.NewRequest("GET","http://10.0.203.112:8888/login/accountLogin.jsp",nil)
	if er!=nil {
		panic(er.Error())
	}
	client :=http.Client{}
	resp,er:=client.Do(req)
	if er!=nil {
		panic(er.Error())
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
