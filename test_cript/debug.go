package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	block, e := aes.NewCipher([]byte("123456789123456789123456"))
	if e!=nil {
		panic(e)
	}
	fmt.Println(block.BlockSize())
}
