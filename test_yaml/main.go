package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func main() {
	type Config struct {
		Host string `yaml:"host"`
		Name string `yaml:"name"`
	}
	var configs = []Config{
		{"10.0.1.1", "db1"}, {"10.0.1.12", "db2"},
	}

	b, e := yaml.Marshal(configs)
	if e!=nil {
		panic(e)
	}
	fmt.Println(string(b))
}
