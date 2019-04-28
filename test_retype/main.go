package main

import (
	"fmt"
	"reflect"
	"strings"
)
type B []byte
type S string
func main() {
	var b = B([]byte("hello"))
	P(b)
	Pr([]byte("hello"))

	//var s = S("world")
	//SS(s)
	//SSr(string("world"))
	var a *int

	fmt.Println(reflect.TypeOf(a).Kind().String())

	arr := strings.Split("", ",")
	fmt.Println(arr[0]=="")
}

func P(b []byte){
	fmt.Println(b)
}

func SS(s string){
	fmt.Println(s)
}

func SSr(s S){
	fmt.Println(s)
}
func Pr(b B){
	fmt.Println(b)
}
