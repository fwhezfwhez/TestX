package main

import (
	"container/list"
	"fmt"
)

func main() {
	type User struct{
		Age int
	}
	user:=User{5}
	l := list.New()
	l.PushFront(user)
	//l.PushBack(1)
	//l.PushBack(2)
	//l.PushBack(3)
	//l.PushBack(4)
	fmt.Println("测试:",l.Front().Value.(User).Age)
	fmt.Println("Before Removing...")
	fmt.Println("长度:",l.Len())
	var n *list.Element
	for e := l.Front(); e != nil; e = n {
		fmt.Println("removing", e.Value.(User).Age)
		n = e.Next()
		l.Remove(e)
	}
	fmt.Println("After Removing...")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}


}