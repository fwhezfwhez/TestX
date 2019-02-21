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
	timestampInt64, _ := strconv.ParseInt("1521421353", 10, 64)
	fmt.Println(timestampInt64)

	str_time := time.Unix(timestampInt64, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str_time)

	//将字符串转换成时间戳

	Time, _ := time.ParseInLocation(startTime, "2018-03-19 09:02:33",time.Local)
	timestamp := Time.Unix()
	fmt.Println(timestamp)
	//fmt.Println("k:",Time)

	Time2, _ := time.ParseInLocation("2006/01/02", "2018/04/19",time.Local)
	fmt.Println(Time2,"333")
	fmt.Println(Time.Unix()<Time2.Unix())
	//时间类型精度修改
	//Local循环不影响, 即可以认为Local和Unix方式存储的空间对象是同一个，不管如何Local Unix都指代一个，不会出现多次-8h
	t3:=Time.Local().Local().Local()
	fmt.Println("Time:",t3.Format(startTime))
	timestamp2:=time.Now().Unix()
	//将时间戳转变成时间
	dataTimeStr1 := time.Unix(timestamp, 0).Local().Format(startTime)
	dataTimeStr2 := time.Unix(timestamp2, 0).Local().Format(startTime)
	fmt.Println(dataTimeStr1)
	fmt.Println("now:",dataTimeStr2)

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

	fmt.Println(time.Now().UnixNano()/1e3)

	var t time.Time = time.Now()
	fmt.Println(t.Format("2006-01-02"))


	heartBeatTick := time.Tick(2 * time.Second)
	for {
		fmt.Println(<-heartBeatTick)
	}
}
