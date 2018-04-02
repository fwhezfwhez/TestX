package main

import (
	"fmt"
	"regexp"
	//"strings"
)

type User struct{
	Id int
}
func main() {

	r, _ := regexp.Compile("^[\u4E00-\u9FA5]*$")
	fmt.Println(r.MatchString("我"))
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。

	//str := "Hello.World"
	//fmt.Println(strings.ContainsAny(str, "."))
	//
	//user :=User{3}
	//user2 :=User{3}
	//fmt.Println(user==user2)
	//
	//fmt.Println(strings.IndexAny("中国麻将IOS", "IOS"))
}