package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main(){
	r:=gin.Default()
	r.GET("/sayHello/",func(c *gin.Context){
		c.JSON(200,"hello")
	})
	r.Run(":8991")
	sql.DB{}.SetConnMaxLifetime()
}
