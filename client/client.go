package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

type Slot struct{
	SlotId string `json:"slot_id,omitempty"`
	SlotName string `json:"slot_name,omitempty"`
}
func main() {
	TestAdd()
	TestUpdate()
	TestQuery()
	TestDelete()
}

func TestAdd(){
	slot:= Slot{SlotName:"广告位2"}
	adddJson,err:=json.Marshal(slot)
	if err!=nil{
		panic(err)
	}
	resp, err := http.Post("http://10.0.203.92:8080/slot/add", "application/json", bytes.NewReader(adddJson))
	defer resp.Body.Close()
	if err!=nil{
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestDelete(){
	slot:= Slot{SlotId:"2"}
	adddJson,err:=json.Marshal(slot)
	if err!=nil{
		panic(err)
	}
	resp, err := http.Post("http://10.0.203.92:8080/slot/delete", "application/json", bytes.NewReader(adddJson))
	defer resp.Body.Close()
	if err!=nil{
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestUpdate(){
	slot:= Slot{SlotId:"3"}
	adddJson,err:=json.Marshal(slot)
	if err!=nil{
		panic(err)
	}
	resp, err := http.Post("http://10.0.203.92:8080/slot/modify", "application/json", bytes.NewReader(adddJson))
	defer resp.Body.Close()
	if err!=nil{
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func TestQuery(){
	resp, err := http.Get("http://10.0.203.92:8080/slots/list")
	defer resp.Body.Close()
	if err!=nil{
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}