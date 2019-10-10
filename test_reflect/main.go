package main

import (
	"encoding/json"
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

    // whether retype
	var a []byte
	var b json.RawMessage
	var c []uint8
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)
	tc := reflect.TypeOf(c)
	fmt.Println(tb.AssignableTo(ta))
	fmt.Println(tc.AssignableTo(ta))
}


type User struct{
	Name string
}

func (u User) XxValidate()(bool,string,error){
	return true, "sucess", nil
}
