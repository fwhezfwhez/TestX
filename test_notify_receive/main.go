package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main(){
	r:= gin.Default()
	r.POST("/get_notify/", getN)
	r.Run(":8111")
}

func getN(c *gin.Context){
	buf ,e :=ioutil.ReadAll(c.Request.Body)
	if e!=nil {
		fmt.Println(e.Error())
		c.String(200, "ok")
		return
	}
	fmt.Println(string(buf))
	c.String(200, "ok")
}
