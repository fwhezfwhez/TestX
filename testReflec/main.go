package main

import (
	"reflect"
	"fmt"
	//"unsafe"
	"encoding/json"
)

type Rule interface{
	InitDefault()
}

type User struct {
	Id   float32
	Name string
	Age string	`Fuck:"kit"`
}
func (u User) InitDefault(){

}
func do(in interface{}){
	_v:= reflect.ValueOf(&in)

	fmt.Println(_v.CanSet())
	_v = _v.Elem()
	fmt.Println(_v.CanSet())
	_v.Field(0).SetString("6")
	fmt.Println(in)
}

func add(a interface{})int{
	var i int
	defer func(int){i=3}(i)
	//fmt.Println(reflect.TypeOf(a))
	return i
}

func t(a ...int){
	fmt.Println(len(a))
}
func main() {
	//t(1,2,3)
	//
	//var a =make([]int,5)
	//a=append(a,6)
	//fmt.Println(a)
	//
	//var k *float64 = nil
	//var t2 interface{}=k
	//fmt.Println("**",t2 == k)
	//fmt.Println("**",k == nil)
	//fmt.Println("**",t2 == nil)
	////fmt.Println(add(""))
	//var id float32= 55
	//example_user := User{id,"ft","32"}
	//
	//fmt.Println(reflect.ValueOf(example_user).Field(0))
	//fmt.Println(reflect.TypeOf(example_user).Field(0).Name)
	//
	////vtype := reflect.TypeOf(example_user)
	//vType :=reflect.TypeOf(1)
	//vValue := reflect.ValueOf(1)
	//fmt.Println(vValue.Field(0).Type().Name)
	//
	//v := vValue.Field(0).Pointer()
	////val:= (*int)unsafe.Pointer(v)
	//val := (*int)(unsafe.Pointer(v))
	//fmt.Println(*val)
	//
	//
	////vType := reflect.TypeOf(1)
	////vValue:=reflect.ValueOf(add)
	////_=vType.NumIn()
	////
	////fmt.Println(vType.String())
	////var argIn = make([]reflect.Value,1)
	////argIn[0] = reflect.ValueOf(1)
	////fmt.Println(vValue.Call(argIn)[0].Int())
	//
	//do(&example_user)
	////_v:= reflect.ValueOf(&example_user)
	////fmt.Println(_v.CanSet())
	////_v = _v.Elem()
	////fmt.Println(_v.CanSet())
	////_v.Field(0).SetString("6")
	////fmt.Println(example_user)
	var a = make(map[string]interface{},0)
	fmt.Println(reflect.TypeOf(a).String())



	//value :=reflect.ValueOf (example_user)
	//fmt.Println(value)
	//
	//des := reflect.Indirect(value)
	//fmt.Println(reflect.TypeOf(des))
	//
	////
	////valueStr :=value.String()
	////fmt.Println(valueStr)
	////
	//v2 :=value.Field(0)
	//fmt.Println(v2)
	//
	//
	//_type :=reflect.TypeOf(example_user)
	//fmt.Println(_type)
	//
	// var tagValue =_type.Field(2).Tag.Get("Fuck")
	// fmt.Println(tagValue)
	//
	// fmt.Println("kk:",_type.Field(2))
	////type_len :=_type.NumField()
	////fmt.Println(type_len)
	////
	////
	////Check(example_user)
	//
	//var subStrings = make([]string,1)
	//subStrings = strings.Split("afdasf",",")
	//fmt.Println(subStrings)
    a["name"] = "ft"
    a["age"] = 9
    buf,er:=json.Marshal(a)
    if er!=nil{
    	fmt.Println(er.Error())
    	return
	}
	type B struct{
		Name interface{} `json:"name"`
		Age interface{} `json:"age"`
	}
	b := B{
		Name: "ft",
	}
	bufb,er:=json.Marshal(b)
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	fmt.Println(string(bufb))
	fmt.Println(string(buf))
	fmt.Println(bufb)
	fmt.Println(buf)
	var b2 = B{}
	er=json.Unmarshal(buf,&b2)
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	var a2 = make(map[string]interface{},0)

	er=json.Unmarshal(buf,&a2)
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	fmt.Println(b2)
}

func Check(input interface{}){
	_type := reflect.TypeOf(input)
	_value :=reflect.ValueOf(input)
	fmt.Println(_type.NumField())
	fmt.Println(_value.Field(0).String())
}
