package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// range sub dir and file
	rd, err := ioutil.ReadDir("G://")
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Println(fi.Name())
		} else {
			fmt.Println(fi.Name())
		}
	}
    fmt.Println(err)

}
