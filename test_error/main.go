package main

import (
	"errors"
	"fmt"
)

type Error struct{ E error }

func (e Error) Error() string {
	return e.E.Error()
}

func f1() error {
	return f2()
}

func f2() error {
	return Error{errors.New("an error")}
}

func main() {
	e := f1()
	switch e.(type) {
	case Error:
		fmt.Println("I am Error")
	case error:
		fmt.Println("I am error")
	}
}
