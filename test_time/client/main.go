package main

import (
	"net/http"
	"fmt"
	//"time"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"time"
)

func main(){
	resp,er := http.Get("http://localhost:8021/time")
	defer resp.Body.Close()
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	var rs interface{}
	bd,er:=ioutil.ReadAll(resp.Body)
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	er = json.Unmarshal(bd,&rs)
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	fmt.Println(reflect.TypeOf(rs))
	fmt.Println(rs)
	t,er:=time.Parse("",rs.(string))
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	fmt.Println(t.UnixNano())
}
