package main

import (
	"runtime"
	"fmt"
	"time"
	"sync"
)

var a int =0

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i:=0;i<100;i++{
		go AddA()
	}
	time.Sleep(10*time.Second)
	//runtime.Gosched()
}
func AddA(){
	mutex :=&sync.Mutex{}
	mutex.Lock()
	a++
	mutex.Unlock()
	fmt.Println(a)
}

