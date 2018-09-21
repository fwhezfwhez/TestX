package main

import (
	"fmt"
	"runtime"
	"sync"
	"os"

	"io/ioutil"
)

func main() {
	var b = make([]byte,16)
	reader,_:=os.Stdin.Read(b[:])
	//for {
	//	_,err:=reader.Read(b[:])
	//	if err!=nil {
	//		if err==io.EOF {
	//			break
	//		}else{
	//			fmt.Println(err.Error())
	//			return
	//		}
	//	}
	//}
	fmt.Println(len(b))
	fmt.Println(b)
	//q1()
	//q2()
	//q3()
	//q4()
	//q5()
	//q6()
	//q7()
	//q8()
	//q9()
	//q10()
}

//考察panic和recover，分析:defer 后设先执行， defer和panic所处的协程不同
func q1() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

//stu是索引,会出现存的3个stu都是一个对象
func q2() {
	type student struct {
		Name string
		Age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println(m["zhou"].Age)
	fmt.Println(m["li"].Age)
	fmt.Println(m["wang"].Age)
}

//先循环完for，出现10，10，10……10，第二个循环是0-9
func q3() {

	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {

			fmt.Println("i: ", i, " 循环1")
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i, " 循环2")
			wg.Done()
		}(i)
	}
	wg.Wait()

}

func q4() {
	var i = 1
	go func(int) {
		fmt.Println("q4:", i)
	}(i)
	i++
	runtime.Gosched()
}
//i索引到了q5 routine的第一个i，输出1
func q5() {
	var i = 1
	go func(i int) {
		fmt.Println("q5:", i)
	}(i)
	i++
	runtime.Gosched()
}
//2
func q6() {
	var i = 1
	go func() {
		fmt.Println("q6:", i)
	}()
	i++
	runtime.Gosched()
}
//子协程形参会索引父域的，使用x即与i无关了，输出2
func q7() {
	var i = 1
	go func(x int) {
		fmt.Println("q7:", i)
	}(i)
	i++
	runtime.Gosched()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}
//略
func q8() {
	t := Teacher{}
	t.ShowA()
}

//select 的唯一和随机
func q9() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
//defer注册和使用的顺序,注册时运行10，20，然后再defer使用2，1
func q10() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
