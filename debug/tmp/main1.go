package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//X-Forwarded-For
	//X-Real-Ip
	// X-Appengine-Remote-Addr
	// X-AppEngine-User-IP
	// X-AppEngine-Remote-Addr
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		fmt.Println(1, context.Request.Header.Get("X-Forwarded-For"))
		fmt.Println(2, context.Request.Header.Get("X-Real-Ip"))
		fmt.Println(3, context.Request.Header.Get("X-Appengine-Remote-Addr"))
		fmt.Println(4, context.Request.RemoteAddr)
		fmt.Println(5, context.Request.Header.Get("X-AppEngine-User-IP"))
		fmt.Println(6, context.Request.Header.Get("X-AppEngine-Remote-Addr"))

		buf,e:=json.MarshalIndent(map[string][]string(context.Request.Header), "  ", "  ")
		if e!=nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println(string(buf))
	})
	r.Run(":8111")
}
