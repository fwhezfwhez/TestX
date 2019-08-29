package main

import (
	"fmt"
	"runtime"
	"time"
)

//每个case语句里必须是一个IO操作
//所有channel表达式都会被求值、所有被发送的表达式都会被求值
//如果任意某个case可以进行，它就执行(其他被忽略)。
//如果有多个case都可以运行，Select会随机公平地选出一个执行(其他不会执行)。
//如果有default子句，case不满足条件时执行该语句。
//如果没有default字句，select将阻塞，直到某个case可以运行；Go不会重新对channel或值进行求值。
//chan广播关闭，只有用写的那方close才行，读方不让close
//break 也是能跳出select的，所以for和select同时存在时，一定要记得 break L
func main() {

    
	runtime.GOMAXPROCS(10)
	var c = make(chan int, 0)
	go routine1(c)
	go routine2(c)
	time.Sleep(1* time.Second)
	close(c)
	time.Sleep(20 * time.Second)
}
func routine1(c chan int) {
	for {
		fmt.Println("r1")
		select {
		case <-c:
			fmt.Println("r1 结束")
			return
		default:
			continue
		}
	}
}

func routine2(c chan int) {
	for {
		fmt.Println("r2")
		select {
		case <-c:
			fmt.Println("r2 结束")
			return
		default:
			continue
		}
	}
}

