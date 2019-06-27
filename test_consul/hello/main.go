package main

import "github.com/gin-gonic/gin"

func main() {
	r:= gin.Default()
	r.GET("/ping/", func(context *gin.Context) {
		context.String(200, "pong")
	})
	r.Run(":7612")
}
