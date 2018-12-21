package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

func main() {
	var in2 = []byte(`{
		"class": {
			"name": "高中1班",
			"master": {
				"name": "张一山",
				"age": 21,
				"company": {
					"name": "go公司",
					"built_by": "张二山",
					"manager": ["张一山", "张二山", "张三山"],
					"country_s": {
						"name": "China",
						"location": "Asure"
					}
				}
			}
		}
	}`)

	var dest map[string]interface{}
	err := json.Unmarshal([]byte(in2), &dest)
	if err != nil {
		log.Fatalf("Unmarshal failed %s", err.Error())
	}
	t := time.Now()
	err = patchStruct(dest, "china2", "class2")
	fmt.Println(time.Now().Sub(t))
	if err != nil {
		log.Fatalf("patch failed %s", err.Error())
	}
	b, err := json.Marshal(dest)
	if err != nil {
		log.Fatalf("Unmarshal failed %s", err.Error())
	}
	fmt.Println(string(b))
	//fmt.Printf("json: %# v", pretty.Formatter(dest))
}

func patchStruct(target interface{}, val interface{}, keys ...string) error {
	return patchSel(reflect.ValueOf(&target), keys, reflect.ValueOf(val))
}

func patchSel(target reflect.Value, sel []string, val reflect.Value) error {
	if len(sel) == 0 {
		return errors.New("invalid sel")
	}
	v := reflect.Indirect(target)
	//v := reflect.ValueOf(target)

Selector:
	switch v.Kind() {
	case reflect.Struct:
		nam := strings.Title(sel[0])
		if len(sel) == 1 {
			f := v.FieldByName(nam)
			if f.IsValid() {

				if f.CanSet() {
					f.Set(val)
					return nil
				}
			}
		} else {
			f := v.FieldByName(nam)
			return patchSel(f, sel[1:], val)
		}
	case reflect.Map:
		// nam := reflect.ValueOf(strings.Title(sel[0]))
		nam := reflect.ValueOf(sel[0])
		if len(sel) == 1 {
			//v.SetMapIndex(nam, val)
			vvvv := v.Interface().(map[string]interface{})
			vvvv[nam.String()] = val.Interface()

			//f := v.MapIndex(nam)
			//if f.IsValid() {
			//	if f.CanSet() {
			//		f.Set(val)
			//		return nil
			//	} else {
			//		f = f.Elem()
			//		f.Set(val)
			//		return nil
			//	}
			//} else {
			//	v.SetMapIndex(nam, val)
			//}
		} else {
			f := v.MapIndex(nam)
			return patchSel(f, sel[1:], val)
		}
	case reflect.Interface:
		vv := v.Interface()
		if vvv, ok := vv.(map[string]interface{}); ok {
			v = reflect.ValueOf(vvv)
			goto Selector
		}
		fallthrough
	default:
		return errors.New("invalid type, must struct/map")
	}
	return nil
}
