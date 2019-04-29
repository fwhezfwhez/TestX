package main

import "fmt"

type User struct {
	Username string
}

func main() {
	var user1  = &User{"ft"}
	var user2 = *user1
	fmt.Println(user2.Username)
	user2.Username = "ft2"
	fmt.Println(user1.Username)

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
