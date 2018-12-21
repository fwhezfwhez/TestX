package main

import "github.com/gin-gonic/gin"

func main(){
	go func(){
		r2 :=gin.Default()
		r2.GET("/hello/", func(c *gin.Context){
			c.JSON(200, "http reply: hello wasm")
		})
		r2.Run(":8890")
	}()

	r :=gin.Default()
	r.GET("/hello/", func(c *gin.Context){
		c.JSON(200, "http reply: hello wasm")
	})
	r.Run(":8889")
}