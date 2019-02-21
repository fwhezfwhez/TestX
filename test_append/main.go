package main

import "fmt"

func main(){
	var a = make([]string ,1, 1)

	a[0] = "1"
	fmt.Printf("%p", a)
	a = append(a, "2")
	fmt.Printf("%p", a)
}
