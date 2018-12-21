package test_obj2map

import (
	"encoding/json"
	"reflect"
	"strings"
)

func Obj2MapByJson(obj interface{}) (map[string]interface{},error){
	var rs =make(map[string]interface{},0)
	buf,er :=json.Marshal(obj)
	if er!=nil {
		return rs,er
	}
	er = json.Unmarshal(buf,&rs)
	if er!=nil {
		return rs,er
	}
	return rs,nil
}

func Obj2MapByReflect(obj interface{})(map[string]interface{}){
	var rs = make(map[string]interface{},0)
	vValue := reflect.ValueOf(obj)
	vType := reflect.TypeOf(obj)
	var tag = ""
	for i:=0;i< vValue.NumField();i++{
		tag = vType.Field(i).Tag.Get("field")
		if tag == "-" {
			continue
		}
		rs[strings.ToLower(vType.Field(i).Name)] = vValue.Field(i).Interface()
	}
	return rs
}