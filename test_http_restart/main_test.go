package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func _case() {
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
					"country": {
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
		fmt.Println(err.Error())
		return
	}
	t := time.Now()
	err = patchStruct(dest, "china2", "class", "master", "company", "country", "name_spec")
	fmt.Println(time.Now().Sub(t))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, err := json.Marshal(dest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(b))
}

func BenchmarkPatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_case()
	}
}


type BigStruct struct {
	C01 int
	C02 int
	C03 int
	C04 int
	C05 int
	C06 int
	C07 int
	C08 int
	C09 int
	C10 int
	C11 int
	C12 int
	C13 int
	C14 int
	C15 int
	C16 int
	C17 int
	C18 int
	C19 int
	C20 int
	C21 int
	C22 int
	C23 int
	C24 int
	C25 int
	C26 int
	C27 int
	C28 int
	C29 int
	C30 int
}

func Loop1(a []*BigStruct) int {
	var n = 0

	for i := 0; i < len(a); i++ {
		n += a[i].C30
	}

	return n
}

func Loop2(a []*BigStruct) int {
	var n = 0

	for _, item := range a {
		n += item.C30
	}

	return n
}

func Loop3(a []BigStruct) int {
	var n = 0

	for i := 0; i < len(a); i++ {
		n += a[i].C30
	}

	return n
}

func Loop4(a []BigStruct) int {
	var n = 0
	var b = make([]BigStruct,len(a))
	copy(b, a)
	for _, item:= range b {
		n += item.C30
	}

	return n
}

func Benchmark_Loop1(b *testing.B) {
	b.StopTimer()
	var a = make([]*BigStruct, 1000)
	for i := 0; i < len(a); i++ {
		a[i] = new(BigStruct)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Loop1(a)
	}
}

func Benchmark_Loop2(b *testing.B) {
	b.StopTimer()
	var a = make([]*BigStruct, 1000)
	for i := 0; i < len(a); i++ {
		a[i] = new(BigStruct)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Loop2(a)
	}
}

func Benchmark_Loop3(b *testing.B) {
	b.StopTimer()
	var a = make([]BigStruct, 1000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Loop3(a)
	}
}

func Benchmark_Loop4(b *testing.B) {
	b.StopTimer()
	var a = make([]BigStruct, 1000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Loop4(a)
	}
}
