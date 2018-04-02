package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"flag"
)

func main() {

	var addr string
	flag.StringVar(&addr,"addr",":8087","example':8087'")

	flag.Parse()
	gin.SetMode(gin.DebugMode)
	router :=gin.Default()
	router.GET("/",CheckToken)

	router.GET("/Generate",TokenHandle)
	router.Run(addr)

}
func TokenHandle(c *gin.Context) {
	c.Writer.Header().Add("x-Auth-Token","1232423423423424")
	c.JSON(200,"哼")
}



func CheckToken(c *gin.Context){
	rq := c.Request

	//fmt.Println(rq)
	header :=rq.Header
	fmt.Println(rq.RemoteAddr)
	addr := header.Get("REMOTE_ADDR")
	fmt.Println("key里的:",addr)
	fmt.Println(header.Get("x-Auth-Token"))
}