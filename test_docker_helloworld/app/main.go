package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main(){
	go func() {
		r1:= gin.Default()
		r1.GET("/sayHello/",func(c *gin.Context){
			c.JSON(200,"hello1")
		})
		r1.Run(":8992")
	}()


	r:=gin.Default()
	r.GET("/sayHello/",func(c *gin.Context){
		c.JSON(200,"hello")
	})
	r.Run(":8991")
}
