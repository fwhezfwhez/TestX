package main

import (
	"net/http"
	"log"
	"fmt"
	_ "net/http/pprof"
	_ "github.com/mkevac/debugcharts"
)

func main() {
	var  a chan int
	go func() {
		fmt.Println("1")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	<-a
}
