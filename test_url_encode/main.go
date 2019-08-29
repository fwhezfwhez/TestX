package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Set("username", "离别时的宁静")
	fmt.Println(v.Encode())

	v,e:=url.ParseQuery("username=%E7%A6%BB%E5%88%AB%E6%97%B6%E7%9A%84%E5%AE%81%E9%9D%99")
	// v,e:=url.ParseQuery("username=%25E7%25A6%25BB%25E5%2588%25AB%25E6%2597%25B6%25E7%259A%2584%25E5%25AE%2581%25E9%259D%2599")

	if e!=nil {
		panic(e)
	}
	fmt.Println(v.Get("username"))
}
