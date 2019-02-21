package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		fmt.Println(c.DefaultQuery("condition",""))
		c.JSON(200, "success")
	})
	r.Run(":8892")
}
