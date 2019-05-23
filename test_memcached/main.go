package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	//"time"
)

func main() {

	mc := memcache.New("localhost:11211", "localhost:11212", "localhost:11213")
	//mc := memcache.New("localhost:11211")
	e := mc.Set(&memcache.Item{Key: "foo2", Value: []byte("3")})
	// 10 second expired
	//e := mc.Set(&memcache.Item{Key: "foo2", Value: []byte("3"), Expiration: 10})
	if e != nil {
		panic(e)
	}
	it, e := mc.Get("foo2")
	if e != nil {
		panic(e)
	}
	//newV,e :=mc.Increment("foo2", 15)
	//if e!=nil {
	//	panic(e)
	//}
	fmt.Println(string(it.Value))

	//fmt.Println(newV)
	//time.Sleep(5 * time.Second)
	//it2, e2 := mc.Get("foo2")
	//fmt.Println(it2, e2)
	//
	//time.Sleep(6 * time.Second)
	//it3, e3 := mc.Get("foo2")
	//fmt.Println(it3, e3)

}
