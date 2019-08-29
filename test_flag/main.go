package main

import (
	"flag"
	"fmt"
)

var mode string
var mode2 string

type PJ struct {
	O string
}

var pj PJ


func main() {
	flag.StringVar(&pj.O, "o", "", "-o rename")
	flag.Parse()
	fmt.Println(pj.O)
}
