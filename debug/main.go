package main

import "fmt"

func main() {
	var m = []int{1,2,3}
	f := func(){
		m = append(m, 4)
	}

	fmt.Println(m)
	f()
	fmt.Println(m)
}
