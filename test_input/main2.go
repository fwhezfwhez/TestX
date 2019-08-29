package main

import (
	"fmt"
)

func scan_demo() {
	var apple_str, orange_str string
	fmt.Scan(&apple_str, &orange_str)
	fmt.Println(apple_str)
	fmt.Println(orange_str)
}
func scanf_demo() {
	var apple_count int
	var apple_name string
	var apple_price float64
	fmt.Scanf("%d %s %f", &apple_count, &apple_name, &apple_price)
	fmt.Println(apple_count)
	fmt.Println(apple_name)
	fmt.Println(apple_price)
}
func scanln_demo() {
	var app_line string
	fmt.Scanln(&app_line)
	fmt.Println(app_line)
}
func main() {
	//scan_demo()
	//scanf_demo()
	scanln_demo()
}
