package main

import (
"fmt"
	"runtime"
	"path"
)

func main() {
	str1:=Pa()
	fmt.Println(str1)
}
func Pa() string{
	_,file,line,_:=runtime.Caller(1)
	fmt.Println(file,line)
	dir :=path.Dir(file)
	return dir
}