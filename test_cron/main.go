package main

import (
	"fmt"
	"github.com/robfig/cron"
)

type TestJob struct {
}

func (this TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("testJob2...")
}

//启动多个任务
func main() {
	c := cron.New()

	//AddFunc
	//秒，分，时，日，月，星
	spec := "0 55 16 * * ?"
	//AddJob方法
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})
	c.AddFunc(spec, func() { fmt.Println(A("1", "2")) })

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()
	select {}
}

func A(a, b string) int {
	return 0
}
