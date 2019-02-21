package main

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)
func main(){
	config.Load(file.NewSource(file.WithPath("../conf.json")))
	fmt.Println(config.Get("hosts").String(""))
}
