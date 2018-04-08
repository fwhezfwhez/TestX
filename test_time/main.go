package main

import (
	"strconv"
	"time"
	"fmt"
)

const (
	startTime = "2006-01-02 15:04:05" //时间起点，固定
)
func main() {

	//将string转换成int64
	timestampInt64, _ := strconv.ParseInt("1448333419", 10, 64)
	fmt.Println(timestampInt64)

	//将字符串转换成时间戳
	Time, _ := time.ParseInLocation(startTime, "2018-03-19 14:55:44",time.Local)
	timestamp := Time.Unix()
	fmt.Println(timestamp)

	timestamp2:=time.Now().Unix()
	//将时间戳转变成时间
	dataTimeStr1 := time.Unix(timestamp, 0).Local().Format(startTime)
	dataTimeStr2 := time.Unix(timestamp2, 0).Local().Format(startTime)
	fmt.Println(dataTimeStr1)
	fmt.Println(dataTimeStr2)

	//前一天
	d,_:=time.ParseDuration("-24h")
	fmt.Println(time.Now().Add(d).Format("2006-01-02"))

	m:=make(map[int]int,3)
	//m[1]=5
	fmt.Println(len(m))
	//m[2]=8
	//m[3]=9
	//m[4]=0
	//fmt.Println(cap(m))
}
