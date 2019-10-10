package main

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		buf, e := ioutil.ReadAll(c.Request.Body)
		if e != nil {
			fmt.Println(errorx.Wrap(e).Error())
			return
		}
		fmt.Println("Recv:", string(buf))
	})
	r.Run(":9191")
}
