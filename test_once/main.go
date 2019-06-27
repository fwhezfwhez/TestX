package main

import (
	"fmt"
	"sync"
)

func main() {
	ite()
}

func ite() {
	var money = 500
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		once.Do(func() {
			money ++
		})
	}
	fmt.Println(money)
}
