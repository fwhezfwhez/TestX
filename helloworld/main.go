package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	s := &http.Server{
		Addr:           ":1111",
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}

	go func(){
		fmt.Println(s.ListenAndServe())

	}()
	go func() {
		fmt.Println(s.ListenAndServe())
	}()
	select {
	}
}
