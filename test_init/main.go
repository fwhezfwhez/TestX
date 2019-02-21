package main

import (
	"fmt"
	"test_X/test_init/p1"
	"test_X/test_init/p2"
)

func init() {
	fmt.Println(3)
}
func main() {
	p1.P1F()
	p2.P2F()
    fmt.Println(6)
}
