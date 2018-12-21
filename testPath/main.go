package main

import (
	"fmt"
	"runtime"
	"path"
)

func main() {
	str1 := Pa()
	fmt.Println(str1)

	type Error struct {
		// basic
		E           error
		StackTraces []string

		// upper
		ReGenerated bool
		Errors      []error
	}
	fmt.Println(Error{
		E:           nil,
		StackTraces: make([]string, 0, 30),
		ReGenerated: false,
		Errors:      make([]error, 0, 30),
	})
}
func Pa() string {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(file, line)
	dir := path.Dir(file)
	return dir
}
