package main

import (
	"log"
	_ "test_X/testLog/case"
)

func main() {
	//file, err := os.OpenFile("./testLog/test.log",  os.O_CREATE | os.O_WRONLY | os.O_APPEND,os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
	//log.SetFlags(0)
	//logger := log.New(file, "", log.LstdFlags|log.Llongfile)
	//logger.Println("日志1.")
	//logger.Println("日志23")
	//
	//var p =make ([]string,0,10)
	//fmt.Println(fmt.Sprintf("%p",p))
	//p=append(p, "a")
	//fmt.Println(fmt.Sprintf("%p",p))
	//p=append(p, "b")
	//fmt.Println(fmt.Sprintf("%p",p))
	//fmt.Println(p)

	log.Println("hello")
	log.Println("%s", "hello2")
	var t = map[string]string{
		"1": "t",
	}
	delete (t,"1")
	delete (t,"1")
}
