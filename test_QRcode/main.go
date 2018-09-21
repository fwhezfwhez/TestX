package main

import (
	 "github.com/skip2/go-qrcode"
	"fmt"
)

func main() {
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err!=nil{
		panic(err)
	}
	fmt.Println(png)
}
