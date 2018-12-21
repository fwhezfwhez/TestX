package main

import (
	"fmt"
	"math"
	"strings"
)
type U struct{}

func main() {
	var a  = make(map[string]*U)
	fmt.Println(a["22"] == nil)

	var str = "^^$$"
	fmt.Println(str[1:len(str)-1])
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

	fmt.Println(RemovePrefix("ax", "axx"))

}

func RemovePrefix(s string,prefix string) string{
	if !strings.HasPrefix(s, prefix){
		return s
	}
	return s[len(prefix):]
}

func Fo(){
	// 0 0 1 2 3 4 5
	// 1 0 1 2 3 4 5
	L:
	for i:=0;i<2;i++{
		fmt.Println("i:",i)
		for j:=0;j<10;j++{
			fmt.Println("j:",j)
			if j ==5 {
				continue L
			}
		}
	}
}
