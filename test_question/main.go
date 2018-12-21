package main

import (
	"fmt"
)

func main() {
	// 结果集
	var rs = make([]string, 0)

	// bit1，bit2，bit3，bit4，bit5 表示万，千，百，十，个位
	for bit1 := 1; bit1 < 10; bit1++ {
		for bit2 := 0; bit2 < 10; bit2++ {
			for bit3 := 0; bit3 < 10; bit3++ {
				for bit4 := 0; bit4 < 10; bit4++ {
					for bit5 := 0; bit5 < 10; bit5++ {
						handle(&rs, []int{bit1, bit2, bit3, bit4, bit5})
					}
				}
			}
		}
	}
	fmt.Println("符合条件的样本数量:",len(rs))
	//打印每条
	fmt.Println("前100条:")
	for i,v:=range rs{
		fmt.Println(v)
		if i==100{
			return
		}
	}
}

// 处理单个数字
// 19382 -> n= []int{1,9,3,8,2}
func handle(rs *[]string, n []int) {
	// 该数字各位之和
	var sum int
	// 剩余2位之和， 该和除以10= result，余数为mod
	var left2, result, mod int
	for _, v := range n {
		sum += v
	}
	// 五位数字中，任意出现某三位和 % 10 =0
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			for k := i + 2; k < 5; k++ {
				if (n[i]+n[j]+n[k])%10 == 0 {
					left2 = sum - n[i] - n[j] - n[k]
					result = left2 / 10
					mod = left2 % 10
					// 将符合的结果，组装进结果集
					*rs = append(*rs, toString(n, []int{n[i],n[j],n[k]},result, mod))
					return
				}
			}
		}
	}
}

// 每条结果集格式
// 99928,其中(9+9+2)%10 =0 ,剩下俩 除以10 等于 1 余7
func toString(n []int, m []int, result int, mod int) string {
	return fmt.Sprintf("%d%d%d%d%d,其中(%d+%d+%d)模10=0,剩下俩除以10等于%d余%d", n[0], n[1], n[2], n[3], n[4],m[0],m[1],m[2], result, mod)
}
