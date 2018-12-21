package main

import (
	"fmt"
)

const (
	Ldatetime = 1 <<iota
	LcausedBy
	Ltrace
)

func main() {
	var flag = Ldatetime | LcausedBy | Ltrace
	Print(flag)
}

func Print(flag int){
	if flag & Ldatetime !=0 {
		fmt.Println("设置打印日期")
	}
	if flag & LcausedBy !=0 {
		fmt.Println("设置打印了原因")
	}
	if flag & Ltrace !=0 {
		fmt.Println("设置打印了堆栈")
	}
}
