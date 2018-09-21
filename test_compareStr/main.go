package main

import (
	"fmt"
	"runtime"
)

func main() {
	var a = 5
	go t(a)
	a =7
	runtime.Gosched()
}


func t(i int){
	fmt.Println(i)
}