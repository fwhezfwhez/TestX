package main

import (
	"reflect"
	"fmt"
	"strings"
)

type User struct {
	Id   string
	Name string
	Age string	`Fuck:"kit"`
}

func main() {
	user := User{"5","ft","32"}
	fmt.Println(user)


	value :=reflect.ValueOf (user)
	fmt.Println(value)

	des := reflect.Indirect(value)
	fmt.Println(reflect.TypeOf(des))

	//
	//valueStr :=value.String()
	//fmt.Println(valueStr)
	//
	//v2 :=value.Field(0)
	//fmt.Println(v2)
	//
	_type :=reflect.TypeOf(user)
	fmt.Println(_type)

	 var tagValue =_type.Field(2).Tag.Get("Fuck")
	 fmt.Println(tagValue)
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
