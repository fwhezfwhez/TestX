package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r.(error).Error())
		}
	}()
	F1()
	F2()
}

func F1() {
	fmt.Println("f1")
	panic(errors.New("111"))
}

func F2() {
	fmt.Println("f2")
}
