package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	var b =5
	fun()(&b)
	fmt.Println(b)
}

func fun() func(*int){
	return changeA
	var c = &gin.Context{}
	c.ClientIP()
}

func changeA(a *int){
	*a++
}
