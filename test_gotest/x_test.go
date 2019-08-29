package test_gotest

import (
	"fmt"
	"testing"
)

func TestX(t *testing.T) {
	fmt.Println(1)

}

// 示例
func ExampleX() {
	fmt.Println(55)
	fmt.Println(5)
	// Output: 55
	// 5
}

// 并行
func BenchmarkX(b *testing.B) {
	b.ReportAllocs()
	b.SetParallelism(3)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Println("hello world" )
		}
	})
}
