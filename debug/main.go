package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type User struct {
	Username string `validate:"regex,^[\u4E00-\u9FA5a-zA-Z0-9_.]{0,40}$"`
}

func (u User) Validate() (bool, string, error) {
	re, e := regexp.Compile("^[\u4E00-\u9FA5a-zA-Z0-9_.]{0,40}$")
	if e != nil {
		return false, "fail becaues " + e.Error(), e
	}
	ok := re.MatchString(u.Username)
	if !ok {
		return ok, "username fails '^[\u4E00-\u9FA5a-zA-Z0-9_.]{0,40}$' got " + u.Username, nil
	}
	return ok, "success", nil
}

func ValidateMethods(input interface{}) (bool, string, error) {
	vType := reflect.TypeOf(input)
	vValue := reflect.ValueOf(input)
	var info string
	var methodName string

	var results []reflect.Value

	for i := 0; i < vType.NumMethod(); i++ {
		methodName = vType.Method(i).Name

		// UserValidate,UserSVValidate
		if strings.HasSuffix(methodName, "Validate"){
			// all cases will validate methods end with 'Validate' or 'SVValidate'
			results = vValue.Method(i).Call(nil)
			if len(results) != 3 {
				info = fmt.Sprintf("while validating method[%d],named '%s',illegal return values,want 3(bool,string,error) but got %d(%s)", i, methodName, len(results), valueListByType(results))
				return false, info, errors.New(info)
			}
			var er error
			ok, msg := results[0].Bool(), results[1].String()
			if results[2].IsNil() {
				er = nil
			} else {
				er = results[2].Interface().(error)
			}
			if ok {
				continue
			} else {
				return ok, msg, er
			}
		}
	}
	return true, "success", nil
}

// when input a []reflect.Value{false, 5, 'example'}
// returns 'bool,int,string'
func valueListByType(r []reflect.Value) string {
	typs := make([]string, 0)
	for _, v := range r {
		typs = append(typs, reflect.TypeOf(v.Interface()).String())
	}
	return strings.Join(typs, ",")
}

func main() {
	u := User{Username: "ft123"}
	ok, msg, e := ValidateMethods(u)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	if !ok {
		fmt.Println(msg)
		return
	}
	fmt.Println(ok, msg)
}
