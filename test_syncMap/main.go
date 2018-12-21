package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 安全map
var sm sync.Map
// 原生map
var m map[interface{}]interface{}
// 原生map锁
var mlock sync.RWMutex
func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	sm = sync.Map{}
	m = make(map[interface{}]interface{})
}
func main(){
	fmt.Println(sm.Load("ok"))
}

func SyncMap(key interface{}, value interface{}, wg *sync.WaitGroup){
	defer wg.Done()
	sm.Store(key, value)
	fmt.Println(sm.Load(key))
}

func Map(key interface{}, value interface{}, wg *sync.WaitGroup){
	defer wg.Done()
	mlock.Lock()

	m[key] = value
    mlock.Unlock()

	mlock.RLock()

	fmt.Println(m[key])

	mlock.RUnlock()
}
