package main

import (
	"reflect"
	"fmt"
	"strings"
)

type Rule interface{
	InitDefault()
}

type User struct {
	Id   string
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
func main() {
	user := User{"5","ft","32"}

	do(&user)
	//_v:= reflect.ValueOf(&user)
	//fmt.Println(_v.CanSet())
	//_v = _v.Elem()
	//fmt.Println(_v.CanSet())
	//_v.Field(0).SetString("6")
	//fmt.Println(user)




	value :=reflect.ValueOf (user)
	fmt.Println(value)

	des := reflect.Indirect(value)
	fmt.Println(reflect.TypeOf(des))

	//
	//valueStr :=value.String()
	//fmt.Println(valueStr)
	//
	v2 :=value.Field(0)
	fmt.Println(v2)


	_type :=reflect.TypeOf(user)
	fmt.Println(_type)

	 var tagValue =_type.Field(2).Tag.Get("Fuck")
	 fmt.Println(tagValue)

	 fmt.Println("kk:",_type.Field(2))
	//type_len :=_type.NumField()
	//fmt.Println(type_len)
	//
	//
	//Check(user)

	var subStrings = make([]string,1)
	subStrings = strings.Split("afdasf",",")
	fmt.Println(subStrings)

}

func Check(input interface{}){
	_type := reflect.TypeOf(input)
	_value :=reflect.ValueOf(input)
	fmt.Println(_type.NumField())
	fmt.Println(_value.Field(0).String())
}
