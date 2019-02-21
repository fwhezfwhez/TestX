package main

import (
	"fmt"
	"regexp"
)

func main() {
	re ,e:= regexp.Compile("^[\u4e00-\u9fa5a-zA-Z0-9_.]{0,40}$")
	if e!=nil {
		panic(e)
	}
	fmt.Println(re.MatchString("ft123"))
}
