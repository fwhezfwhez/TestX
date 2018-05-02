package main

import (
"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	router :=gin.Default()
	router.GET("/",T)

	//用crt和key
	//er:=http.ListenAndServeTLS(":8081", "server.crt",
	//	"server.key", nil)
	//if er!=nil {
	//	fmt.Println(er)
	//}

	//
	err := http.ListenAndServe(":8080",cors.AllowAll().Handler(router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func T(c *gin.Context){
	fmt.Fprintf(c.Writer,
		"Hi, This is an example of https service in golang!")
	fmt.Println("hello")
	c.JSON(1,"hello")
}