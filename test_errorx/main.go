package main

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
)

func main() {
	e1 := f1()
	e2 := f2()

	fmt.Println(errorx.GroupErrors(errorx.Wrap(e1), errorx.Wrap(e2)))
}

func f1() error {
	return errorx.NewFromString("err1")
}

func f2() error {
	return errorx.NewFromString("err1")
}
