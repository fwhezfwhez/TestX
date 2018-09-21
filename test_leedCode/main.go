package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var arr = []int{
		1,5,9,8,12,100,2,7,3,3,1,94,96,93,98,99,97,
	}
	fmt.Println(MaxSerialArr(arr))
}

//1. 两数之和
func twoSum(nums []int, target int) []int {
	var len1 = len(nums)
	for i := 0; i < len1-1; i++ {
		for j := i + 1; j < len1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//x.1 求数组里最长连续字数组 O(n)内
func MaxSerialArr(arr []int) []int {
	//总长
	var max = 0
	//集合点
	var jointIndex = -1
	//结果集
	var rs = make([]int, 0, len(arr))
	//集合点左边(比集合点小的部分长度)
	leftLen:= 0
	for i, v := range arr {
		lengthTmp,leftLenTmp,_:= Depth(arr,v)
		if lengthTmp > max {
			max = lengthTmp
			jointIndex = i
			leftLen = leftLenTmp
		}
	}
	min := arr[jointIndex] - leftLen
	for j:=0;j<max;j++{
		rs = append(rs,min+j)
	}
	return rs
}

//总深度，左深度，右深度
func Depth(arr []int, elem int) (int,int,int) {
	return  1 + containLeftLen(arr, elem-1) + containRightLen(arr, elem+1),containLeftLen(arr,elem-1), containRightLen(arr,elem+1)
}
//arr是否包含elem
func contain(arr []int, elem int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
//左深度
func containLeftLen(arr []int, elem int) int {
	if contain(arr, elem) {
		return containLeftLen(arr, elem-1)
	}
	return 0
}
//右深度
func containRightLen(arr []int, elem int) int {
	if contain(arr, elem) {
		return containRightLen(arr, elem+1) + 1
	}
	return 0
}
