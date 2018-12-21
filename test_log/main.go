package main

import (
	"log"
	"os"
	_ "test_X/test_log/case"
)

func main() {
	file, err := os.OpenFile("G:/go_workspace/GOPATH/src/test_X/testLog/test.log",  os.O_CREATE | os.O_WRONLY | os.O_APPEND,os.ModePerm)
	if err != nil {
		panic(err)
	}
	//log.SetFlags(0)
	logger := log.New(file, "", log.LstdFlags|log.Llongfile)
	logger.SetPrefix("[id=10101011] ")
	logger.Println("测试")


	//var p =make ([]string,0,10)
	//fmt.Println(fmt.Sprintf("%p",p))
	//p=append(p, "a")
	//fmt.Println(fmt.Sprintf("%p",p))
	//p=append(p, "b")
	//fmt.Println(fmt.Sprintf("%p",p))
	//fmt.Println(p)

	//log.Println("hello")
	//log.Println("%s", "hello2")
	//var t = map[string]string{
	//	"1": "t",
	//}
	//delete (t,"1")
	//delete (t,"1")
	//
	//var k = make(map[int]int)
	//k[6] ++
	//fmt.Println(k[6])
}
