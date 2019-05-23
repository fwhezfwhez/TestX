package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"utils"
)

func main() {
	r:= gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(400, "ok")
		context.Abort()
		fmt.Println(context.Writer.Status())
		})
	r.Run(":7567")
	utils.IfZero()
}
