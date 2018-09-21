package main

import "fmt"

type String string

func main() {
	a := "5"
	b := String(a)

	fmt.Println(string(b))
	ad(string(b))
}

func ad(a string){
	fmt.Println(1)
}