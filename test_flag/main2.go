package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type U struct {
	Stream   []byte
	State    *int
	Handlers []func()
	L        *sync.RWMutex
	title    string
}

func (u *U) x() {
	u.L.Lock()
	defer u.L.Unlock()
	fmt.Println(u.title, 444)
	time.Sleep(2 * time.Second)
}

func (u *U) y() {
	u.L.Lock()
	defer u.L.Unlock()
	fmt.Println(u.title, 555)
	time.Sleep(2 * time.Second)
}
func main() {
	sync.RWMutex{}
	runtime.GOMAXPROCS(runtime.NumCPU())
	var a = 2
	var hs = make([]func(), 0, 10)
	hs = append(hs, func() {
		fmt.Println(1)
	})
	var u = U{
	}
	fmt.Println(u)

	u2 := U{
		u.Stream,
		u.State,
		u.Handlers,
		u.L,
		"u2",
	}
	u2.Stream = []byte{5}
	fmt.Println(u.Stream)
	*(u2.State) = 3
	fmt.Println(*(u.State))
	u2.Handlers = append(u2.Handlers, func() {
		fmt.Println(2)
	})
	fmt.Println(len(u.Handlers))

	for i := 0; i < 100; i++ {
		go u.x()
		go u2.x()
		go u.y()
		go u2.y()
	}

	select {}
}
