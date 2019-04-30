package main

import (
	"fmt"
	"strings"
)

type User struct {
	Username string
	O        *Object
}
type Object struct {
	ObjectName string
}

func DoUser(u User) {
	u.O.ObjectName = "PP"
}
func main() {
	fmt.Println(f("0d0d"))
}

func f(str string) string {
	if len(str) %2 !=0 {
		panic("str length should %2 =0")
	}
	var tmp = make([]string, 0, len(str)/2)

	for i:=0;i<len(str);i++{
		if i %2 ==1 {
			tmp = append(tmp, string(str[i-1:i+1]))
		}
	}
	return strings.Join(tmp, "\\x")
}
func B(f func()) {}

type U struct{}

func (u U) A() {

}
func 交集(a, b []int) []int {
	var rs = make([]int, 0)
	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				rs = append(rs, v)
			}
		}
	}
	return rs
}

func 差集(a, b []int) []int {
	var rs = make([]int, 0)
	for _, v := range a {
		if !In(v, b) {
			rs = append(rs, v)
		}
	}
	return rs
}

// s 是否in arr
func In(s int, arr []int) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
