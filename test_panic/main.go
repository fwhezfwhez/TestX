package main

import (
	"fmt"
	"time"
)

func main() {
    defer func(){
    	fmt.Println(2)
	}()
	defer func(){
		fmt.Println(1)
	}()
	go func (){
		defer func(){
			fmt.Println(2)
		}()
		defer func(){
			fmt.Println(1)
		}()
		panic(0)
	}()
    time.Sleep(5 * time.Second)
}

