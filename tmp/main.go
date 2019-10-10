package main

import (
	"fmt"
)

func main() {
	var nums = make([]int, 2,10)
	for i, v := range nums {
		if i == 0 {
			nums[i+1]++
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println(nums)
}
