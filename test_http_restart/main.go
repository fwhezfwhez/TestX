package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	r := gin.Default()
	r.GET("/r1/", func(c *gin.Context) {
		fmt.Println("r1")
		c.JSON(200, "r1")
	})

	s := &http.Server{
		// :8112默认
		Addr:           ":8888",
		Handler:        cors.AllowAll().Handler(r),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}
	go func(){
		if e:=s.ListenAndServe();e!=nil{
			fmt.Println(e.Error())
		}
	}()

	go func(){
		time.Sleep(10 * time.Second)
		fmt.Println("10s pass")
	    if e:=s.Shutdown(context.Background());e!=nil{
	    	panic(e)
		}
	}()

	go func(){
		time.Sleep(20 * time.Second)
		fmt.Println("20s pass")
		r.GET("/r2/", func(c *gin.Context) {
			fmt.Println("r2")
			c.JSON(200,"r2")
		})
		s.ListenAndServe()
	}()

	select{}
}
