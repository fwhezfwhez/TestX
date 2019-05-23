package main

import (
	"fmt"
	_ "github.com/mkevac/debugcharts"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	var a chan int
	go func() {
		fmt.Println("1")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	<-a
}
