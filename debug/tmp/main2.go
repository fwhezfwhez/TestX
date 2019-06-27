package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println(errors.Wrap(fmt.Errorf("hello error"), "hapyy"))
}
