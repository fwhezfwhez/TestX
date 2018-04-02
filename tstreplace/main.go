package main

import (
	"fmt"
	"log"
	"runtime"
)

type Error struct {
	Err error
	Location string
	Advice string
}
func ErrorWrap(err error,advise string) *Error{

	_,file,line,_:=runtime.Caller(1)
	location := fmt.Sprintf("%s %d",file,line)
	errTemp := new(Error)
	errTemp.Err = err
	errTemp.Location = location
	errTemp.Advice = advise
	return errTemp
}
func LogPrintln(err error,advise string){
	log.Println(ErrorWrap(err,advise).String())
}

func (err *Error) String() string{
	return fmt.Sprintf("msg:%v,location:%s,advice:%s",err.Err,err.Location,err.Advice)
}

func main(){
	var a = make(map[string]interface{})
	aq,ok:=a["d"]
	fmt.Println(aq,ok)
	 var k bool
	var sting =make([]string,0)
	sting = append(sting,"a")
	sting = append(sting,"b")
	T(sting...)
	fmt.Println(k)

	var ss =make([]string,0)
	ss=append(ss,"kkk")
	ss=append(ss,"kkk2")
	fmt.Println(ss)
	ss=remove(ss,"kkk22")
	fmt.Println(ss)

	var kd =make([]string,0,3000)
	kd=append(kd,"333")
	fmt.Println(cap(kd),len(kd))
	kd = make([]string,0,3000)
	fmt.Println(cap(kd),len(kd))

	i:=make(map[string]map[string]int)
	ii:=make(map[string]int)
	ii["b"]=5
	i["a"]=ii
	fmt.Println(i)
}

func T(a ...string){
	fmt.Println(a)
}

func remove(slice []string, elems string) []string {
	for i,v := range slice {
		if v == elems {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return  slice
}

