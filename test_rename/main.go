package main

import (
	"fmt"
)

type String string

func main() {
	//a := "5"
	//b := String(a)
	//
	//fmt.Println(string(b))
	//ad(string(b))
	//ab(a)
	c := make(chan int ,0)
	go func(ch chan int){
		for{
			fmt.Println(1)
			select{
			    case <-ch:
			    	fmt.Println("exit success")
			    	break
			    default:
			    	continue
			}
		}
	}(c)
	e:=ch(c)
	if e!=nil {
		fmt.Println(e.Error())
	}
}

func ad(a string){
	fmt.Println(1)
}
func ab(b String){
	fmt.Println(2)
}

func ch( c chan int)error{
	select {
	case c <- 1:
		fmt.Println("close handle")

	}
	return nil
}