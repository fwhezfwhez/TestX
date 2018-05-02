package main

import (
	"fmt"
	"math"
)

func main() {
	//1.表示一个非number值
	fmt.Println(math.NaN()==math.NaN())

	//2.表示一个无穷的大或者无穷小
	fmt.Println(math.Inf(1),math.Inf(-1))
	fmt.Println(1/(2./3/math.Inf(-1)))

	//3.二次方跟
	fmt.Println(math.Sqrt(9.))

	//4.三次方根
	fmt.Println(math.Cbrt(27))

	fmt.Println(math.Sin(math.Pi))

	fmt.Println(5/2)
	fmt.Println(math.Pow(4,2))
}
