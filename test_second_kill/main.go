package main

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	"runtime"
	"sync"
	"time"
)

var GoodsA = SecondObject{&sync.RWMutex{}, 10}

func main() {
	// 单例消费
	fmt.Println("案例一: 单笔消费")
	GoodsA.SecondKill(5*time.Second, 3, func(r chan Result) {
		// 支付订单
		time.Sleep(4 * time.Second)
		r <- Result{1, nil}
	})
	fmt.Println("消费成功，剩余:", GoodsA.Num)

	// 超时消费
	fmt.Println("案例2: 消费超时")
	n, token, e := GoodsA.SecondKill(3*time.Second, 3, func(r chan Result) {
		// 支付订单
		time.Sleep(4 * time.Second)
		r <- Result{1, nil}
	})
	fmt.Println(n, token, e)

	// 库存不足
	fmt.Println("案例3: 库存不足")
	n,token,  e = GoodsA.SecondKill(3*time.Second, 11)
	fmt.Println(n,token, e)

	fmt.Println("案例4: 1000个请求，抢购剩余7个")
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			defer wg.Done()
			n, token,  e := GoodsA.SecondKill(5*time.Second, 1)
			fmt.Println(fmt.Sprintf("request %d result:", i), n, token, e)
		}(i)
	}
	wg.Wait()

}

type SecondObject struct {
	M   *sync.RWMutex
	Num int
}

// 下单到支付成功之间的handler result
// 1,nil 成功
// 2, error 错误
type Result struct {
	Status int
	Err    error
}

// 0,nil success
// 500, error 系统错误
// 400, error 请求错误
// 2, error 库存不足，卖完了
// 3, error 购买数量大于库存
// 4, 超时
func (s *SecondObject) SecondKill(timeout time.Duration, buyNum int, handlers ... func(results chan Result)) (int, interface{}, error) {

	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()
	s.M.Lock()
	defer s.M.Unlock()
	if buyNum > s.Num {
		return 3, nil, errors.New(fmt.Sprintf("buy number %d more than saved number %d", buyNum, s.Num))
	}
	if s.Num < 0 {
		return 2, nil, errors.New("save out")
	}

	var jobNum = len(handlers)
	result := make(chan Result)

	for _, v := range handlers {
		go v(result)
	}

	for {
		if jobNum == 0 {
			break
		}
		select {
		case <-time.After(timeout):
			return 4, nil, errors.New("time out")
		case jobResult := <-result:
			switch jobResult.Status {
			case 1:
				jobNum --
			case 2:
				return jobResult.Status, nil, jobResult.Err
			}
		}
	}

	s.Num -= buyNum
	// 每单的票据
	// 如果需要自定义票据, 删改下两行
	token, _ := uuid.NewV4()
	return 0,token.String(), nil
}
