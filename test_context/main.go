package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// with cancel
	ctx, cancel := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "user_id", "ft")
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("routine1 stops")

				return
			default:
				fmt.Println("1")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("routine2 stops")
				return
			default:
				fmt.Println("2")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx2)

	time.Sleep(5*time.Second)

	cancel()
	time.Sleep(5 * time.Second)
	//go watch(ctx, "【监控1】")
	//go watch(ctx, "【监控2】")
	//go watch(ctx, "【监控3】")
	//time.Sleep(10 * time.Second)
	//fmt.Println("可以了，通知监控停止")
	//cancel()
	////为了检测监控过是否停止，如果没有监控输出，就表示停止了
	//time.Sleep(5 * time.Second)

	// with time out cancel
	// 1. ctx 不共享， 每个任务请使用新的ctx不能复用
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()
	//go func(c context.Context){
	//	for {
	//		fmt.Println("1")
	//		time.Sleep(1 * time.Second)
	//	}
	//}(ctx)
	//time.Sleep(3*time.Second)
	//go func(c context.Context){
	//	for {
	//		fmt.Println("2")
	//		time.Sleep(1 * time.Second)
	//	}
	//}(ctx)
	//select {
	//	case <-ctx.Done():
	//		fmt.Println("超时了")
	//}

	// 2. 执行任务，任务完毕或者超时时结束
	//wg := sync.WaitGroup{}
	//wg.Add(100)
	//ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	//for i := 0; i < 100; i++ {
	//	go func(m int, c context.Context, w *sync.WaitGroup) {
	//		n := 0
	//		for {
	//			n++
	//			fmt.Println(m)
	//			time.Sleep(1 * time.Second)
	//			if n ==9 {
	//				wg.Done()
	//			}
	//		}
	//	}(i, ctx, &wg)
	//}
	//
	//select{
	//	case <-ctx.Done():
	//		fmt.Println("超时了")
	//	case <-wait(&wg):
	//		fmt.Println("执行完毕")
	//}

	//3. withValue 竟然只能放一组值
	//ctx := context.WithValue(context.Background(), "username", 123)
	//fmt.Println(ctx.Value("username"))
}

func wait(wg *sync.WaitGroup) <-chan bool {
	var ch = make(chan bool)
	go func() {
		wg.Wait()
		ch <- true
	}()
	return ch
}
