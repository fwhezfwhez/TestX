package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)

func main(){
	r:=gin.Default()
	r.GET("/time",func(c *gin.Context){
		t:=time.Now()
		fmt.Println("服务端产出的时间:",t.UnixNano())
		c.JSON(200,t)
	})
	r.Run(":8021")
	
}
