package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"flag"
	"encoding/xml"
	"time"
	"net/http"
	"log"
	_ "net/http/pprof"
)

//type User struct{
//	Name string `json:"name"`
//}
func init(){

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
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

	router.GET("/example_user/:uname/:id",TP)
	router.PUT("/example_user/:uname/:id",TP)
	//router.POST("/TP",TP)

	router.DELETE("/example_user/:uname/:id/delete",TP)
	router.GET("/timeout",TT)


	router.Run(addr)

}
func TT(c *gin.Context){
	fmt.Println("in")
	time.Sleep(100*time.Second)
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
func TP(c *gin.Context){
	fmt.Println("in TP")
	//time.Sleep(20*time.Second)

	un:=c.Param("uname")
	id:=c.Param("id")

	fmt.Println("uname",un)
	fmt.Println("id:",id)

	time.Sleep(5*time.Second)
	v:=c.DefaultQuery("ft","")
	fmt.Println("query,ft:",v)
	type U struct{
		Uname string `form:"uname" json:"uname"`
		Id string	`form:"id" json:"id"`
	}
	var u U
	c.Bind(&u)
	fmt.Println(u.Uname,u.Id)
	c.JSON(200,"ok")
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
