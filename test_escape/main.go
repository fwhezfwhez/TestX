package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f()string{
	var tmp string
	for i:=1;i<100;i++{
		tmp += "&"
	}
	return tmp
}

func f2()*string{
	var tmp string
	for i:=1;i<100;i++{
		tmp += "&"
	}
	return &tmp
}
