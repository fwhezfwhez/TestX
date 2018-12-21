package test_obj2map

import (
	"fmt"
	"testing"
	"time"
)

func TestObj2MapByJson(t *testing.T) {
	u := struct{
		Name string `json:"name"`
		Age int `json:"age"`
		Sal float64 `json:"sal"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Name:"ft",
		Age:9,
		Sal:1000.1,
		CreatedAt: time.Now(),
	}
	m,e:=Obj2MapByJson(u)
	if e!=nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(m)
}

func TestObj2MapByReflect(t *testing.T) {
	type k struct{
		Name2 string
	}
	u := struct{
		Name string `json:"name"`
		Age int `json:"age"`
		Sal float64 `json:"sal"`
		CreatedAt time.Time `json:"created_at"`
		K k
	}{
		Name:"ft",
		Age:9,
		Sal:1000.1,
		CreatedAt: time.Now(),
		//K: k{Name2:"k"},
	}
	m:=Obj2MapByReflect(u)

	fmt.Println(m)
}

func BenchmarkObj2MapByJson(b *testing.B) {
	u := struct{
		Name string `json:"name"`
		Age int `json:"age"`
		Sal float64 `json:"sal"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Name:"ft",
		Age:9,
		Sal:1000.1,
		CreatedAt: time.Now(),
	}
	var m map[string]interface{}
	var e error
	for i:=0;i<b.N;i++{
		m,e=Obj2MapByJson(u)
		if e!=nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println(m)
	}
}


func BenchmarkObj2MapByReflect(b *testing.B) {
	u := struct{
		Name string `json:"name"`
		Age int `json:"age"`
		Sal float64 `json:"sal"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Name:"ft",
		Age:9,
		Sal:1000.1,
		CreatedAt: time.Now(),
	}
	var m map[string]interface{}
	for i:=0;i<b.N;i++{
		m=Obj2MapByReflect(u)
		fmt.Println(m)
	}
}