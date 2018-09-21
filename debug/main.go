package main

import (
	"time"
	"runtime"
	"github.com/fwhezfwhez/go-queue"
	"strconv"
	"fmt"
	"sync"
	"strings"
)
var mutex sync.Mutex

func init() {
	fmt.Print(time.Now().UnixNano())
}
func main() {
	fmt.Println(len(strings.Split(":",":")))


	runtime.GOMAXPROCS(runtime.NumCPU())
	wg :=sync.WaitGroup{}


	var queue = Queue.New(5000)
	//var rs = make([]int64,0)
	for i:=0;i<5000;i++ {
		wg.Add(1)
		go func(in int,q *Queue.Queue){
			tmp:= getNum()
			//add(&rs, tmp)
			q.SafePush(tmp)
			fmt.Println("finish "+strconv.Itoa(in)+",produce:"+strconv.FormatInt(tmp,10))
			wg.Done()
		}(i,queue)
	}
	wg.Wait()
	queue.Print()
	fmt.Println(queue.ValidLength())
	fmt.Println(queue.Length())

	//fmt.Println(rs)
	//fmt.Println(len(rs))
}

func getNum()int64{
	mutex.Lock()
	defer mutex.Unlock()
	return time.Now().UnixNano()
}
func add(rs *[]int64,i int64){
	mutex.Lock()
	defer mutex.Unlock()
	*rs = append(*rs,i)
}