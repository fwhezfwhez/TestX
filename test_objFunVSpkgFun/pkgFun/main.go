package main

import (
	"time"
	"test_X/test_objFunVSpkgFun/pkgFun/Dao"
	"fmt"
)
var totalTime time.Duration
var averageTime time.Duration
var timeStart time.Time
var timeEnd time.Time
var marchbenchN int = 10000

type Order struct{
	Id int
	OrderId string `xorm:"orderId"`     //如果不做适配，会默认编程order_id导致数据库找不到
	Cid string
	Pversion string
	Ttime time.Time
	Ip string
	OrderStatus int `xorm:"orderStatus"`
	Phone string
	Uname string
	PayStatus int `xorm:"payStatus"`
}

func main() {
	orders :=make([]Order,0)

	timeStart = time.Now()
	for i:=0;i<marchbenchN;i++{
		Dao.Query(&orders)
		if i!= marchbenchN-1 {
			orders = make([]Order,0)
		}
	}
	timeEnd = time.Now()
	totalTime=timeEnd.Sub(timeStart)
	averageTime = totalTime/time.Duration(marchbenchN)
	fmt.Println("总时间：",totalTime,"平均:",averageTime)
	fmt.Println(orders)
}