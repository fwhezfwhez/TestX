package main

import (
	"flag"
	"fmt"
	"time"
)

var mode string
var mode2 string
func init() {
	flag.StringVar(&mode, "mode", "default", "go run main.go -mode 'dev'")
	flag.StringVar(&mode2, "mode2", "default", "go run main.go -mode2 'dev'")

	flag.Parse()
}

func main() {
	fmt.Println(time.ParseDuration("1h55m"))
	fmt.Println(mode, mode2)

	fmt.Println(fmt.Println(time.Now().Unix()))
}

