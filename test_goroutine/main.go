package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var workNum = 2000
	wg := sync.WaitGroup{}
	wg.Add(workNum)
	WrapTime(func(){
		for i:=0; i<workNum;i++ {
			go RemoteServe(8 * time.Second, &wg)
		}
		wg.Wait()
	})
}


func RemoteServe(cost time.Duration, wg *sync.WaitGroup) string {
	defer wg.Done()
	time.Sleep(cost)
	return "ok"
}

func WrapTime(f func()){
	t1 := time.Now()
	f()
	fmt.Println(time.Now().Sub(t1).String())
}