package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	// 每秒限制执行30次
   //  TimesPerSecond()
   // 跳过一些次数
   //T2()
   // 使用Reverse来获取下一次执行的时间，Reverse + sleep 等价于wait
   T3()
}

func TimesPerSecond() {
	l := rate.NewLimiter(1, 30)

	go func() {
		for {
			l.Wait(context.Background())
			fmt.Println("do")
		}
	}()

	select {
	case <-time.After(1 * time.Second):
		fmt.Print("1s到时")
	}
}

func T2() {
	l := rate.NewLimiter(1, 30)
	go func(){
		for {
			if l.AllowN(time.Now(), 31) {
				fmt.Println("do")
			} else {
				// fmt.Println(time.Now().Format("undo"))
			}
		}
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Print("2s到时")
	}
}

func T3() {
	l := rate.NewLimiter(1, 30)

	go func() {
		for {
			r:=l.Reserve()
			time.Sleep(r.Delay())
			fmt.Println("do")
		}
	}()

	select {
	case <-time.After(1 * time.Second):
		fmt.Print("1s到时")
	}
}
