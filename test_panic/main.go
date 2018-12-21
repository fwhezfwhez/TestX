package main

import "fmt"

func init(){

}

func main() {
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println(r)
		}
	}()
	f()
}

func f(){
	panic(111)
}
