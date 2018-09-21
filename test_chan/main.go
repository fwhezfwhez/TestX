package main

import (
	"fmt"
	"time"
	"bytes"
	"io"
	"io/ioutil"
)


func main(){
	ioutil.ReadAll()

	src :=[]byte("hello")
	src2 :=[]byte("hello2")
	buf :=new(bytes.Buffer)
	buf.Write(src)

	fmt.Println(buf.String())
	buf.Write(src2)
	fmt.Println(buf.String())

	var rs  = make([]byte,buf.Len())
	buf.Read(rs[:5])
	fmt.Println("rs",string(rs))
	buf.Read(rs[5:])
	fmt.Println("rs",string(rs))



	var a = make(chan string,0)


	go func(){
		fmt.Println("1,"+<-a)
	}()
	go func(){
		fmt.Println("2,"+<-a)
	}()
	go func(){
	   a<-"1"
	}()
	time.Sleep(2*time.Second)


}