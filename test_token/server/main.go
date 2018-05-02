package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"flag"
	"encoding/xml"
)

func main() {

	var addr string
	flag.StringVar(&addr, "addr", ":8088", "example':8087'")
	flag.Parse()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/", CheckToken)

	router.GET("/Generate", TokenHandle)
	router.POST("/Test", Test)
	router.POST("/TestBinding",T)
	router.GET("/TestBindQuery",BQ)


	router.Run(addr)

}

type User struct {
	Name string `form:"name" binding:"required"`
	Age string `form:"age"`
}
type Xml struct {
	XMLName          xml.Name `xml:"xml"`
	Age string
	Year int
}
func Test(c *gin.Context) {
	xxml := Xml{}
   c.Bind(&xxml)
   fmt.Println(xxml)
   c.JSON(200,"OK")
}
func TokenHandle(c *gin.Context) {
	c.Writer.Header().Add("x-Auth-Token", "1232423423423424")
	c.JSON(200, "哼")
}

func CheckToken(c *gin.Context) {
	rq := c.Request

	//fmt.Println(rq)
	header := rq.Header
	fmt.Println(rq.RemoteAddr)
	addr := header.Get("REMOTE_ADDR")
	fmt.Println("key里的:", addr)
	fmt.Println(header.Get("x-Auth-Token"))
}


func T(c *gin.Context){
	user :=User{}
	user.Name =  "undefined"
	user.Age = "undefined"
	er:=c.Bind(&user)
	if er!=nil {
		c.JSON(400,er.Error())
	}
}

func BQ(c *gin.Context){
	user :=User{}
	er:=c.BindQuery(&user)
	if er!=nil {
		c.JSON(400,er.Error())
		return
	}
	fmt.Println(user)
}