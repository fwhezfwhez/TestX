package main

import (
	"fmt"
)

func main() {
fmt.Println(int32(int64(4)))
}
//
//func main() {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	var wg = sync.WaitGroup{}
//	for i:=0;i<60000;i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			var j int32 = 400
//			_ = ToInt32(j)
//			_ = ToString(strconv.Itoa(int(j)))
//		}()
//	}
//	wg.Wait()
//}

// 正常情况
func ToInt32(i int32) int32{
	return i
}

func ToInt32Ptr(i int32) *int32{
	return &i
}

func ToStringPtr(i string) *string{
	return &i
}

func ToString(i string) string{
	return i
}
