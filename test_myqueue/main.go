package main

import (
	"fmt"
	queue "github.com/fwhezfwhez/go-queue"
	"time"
)

func main() {
	//初始化,init
	q := queue.NewEmpty()
	//压入,push
	q.Push(5)
	q.Push(4)
	//打印,print
	q.Print()
	//出列,pop
	fmt.Println(q.Pop())
	//打印,print
	q.Print()
	//长度,len
	fmt.Println(q.Length())
	//并发安全压入,currently safe push
	q.SafePush(6)
	//并发安全出列,currently safe pop
	fmt.Print(q.SafePop())
	q.Print()

	// time queue
	tq := queue.TimeQueueWithTimeStep(10*time.Second, 50, 1*time.Nanosecond)
	tq.StartTimeSpying()
	tq.TPush(5)
	tq.SafeTPush(6)

	fmt.Println("init:")
	tq.Print()

	time.Sleep(5 * time.Second)
	fmt.Println("after 5s:")
	tq.Print()

	time.Sleep(9 * time.Second)
	fmt.Println("after 14s")
	tq.Print()
}
