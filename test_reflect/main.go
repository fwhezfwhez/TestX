package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main(){
	str :="UserSVValidateSVBCreate"
	i:=strings.Index(str, "SVdsfaf")
	fmt.Println(i)
	fmt.Println(str[i+len("SVValidate"):])

	var u = User{"ft"}
	vType := reflect.TypeOf(u)
	vValue := reflect.ValueOf(u)

	methods := vType.Method(0)
	fmt.Println(methods.Name)

	//var in = make([]reflect.Value,0)
	results := vValue.Method(0).Call(nil)
	fmt.Println(results)

	r1 := results[0].Interface()
	fmt.Println(reflect.TypeOf(r1).String())
}


type User struct{
	Name string
}

func (u User) XxValidate()(bool,string,error){
	return true, "sucess", nil
}