package main

import "fmt"
type B []byte
type S string
func main() {
	var b = B([]byte("hello"))
	P(b)
	Pr([]byte("hello"))

	//var s = S("world")
	//SS(s)
	//SSr(string("world"))
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