package main

import (
	"errors"
	"fmt"
	"net/http"

	"runtime/debug"
)

func main() {
	rs := http.Response{}
	if rs.Body != nil {
		rs.Body.Close()
	}
	e := makeError()
	if e != nil {
		fmt.Println(e.Error())
		fmt.Println(string(debug.Stack()))
	}
}

func makeError() error {
	fmt.Println("do sth!")
	return errors.New("error happen")
}
