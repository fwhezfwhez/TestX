package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	r.POST("/form/", func(c *gin.Context) {
		type P struct{
			UserID string `form:"user_id"`
			Pwd string `form:"password"`
		}
		var p P
		c.Bind(&p)
		fmt.Println(p)
		c.JSON(200, p)
	})
	r.Run(":7999")
}
