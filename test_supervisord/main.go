package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(200, "ok")
	})
	r.Run(":7111")
}
