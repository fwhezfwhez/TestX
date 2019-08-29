package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()
	r.POST("/1", func(c *gin.Context) {
		fmt.Println(c.Request)
		c.Redirect(304, "http://localhost:8808/2")
	})

	r.POST("/2", func(c *gin.Context) {
		fmt.Println(c.Request)
	})
	r.Run(":8808")
}
