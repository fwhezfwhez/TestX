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
