package main

import (
	"fmt"
)

const (
	Sucess = iota+1
	Fail
	Validating
)
func main() {
	fmt.Println(Sucess)
	fmt.Println(Fail)
	fmt.Println(Validating)
//1 2 3
}