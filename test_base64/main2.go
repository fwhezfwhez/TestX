package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var str = "hahahha"
	sct := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(sct)
	var buf []byte
	buf,e:=base64.StdEncoding.DecodeString("1df09d0a1677dd72b8325aec59576e0c")
	fmt.Println(string(buf), len(buf),e)
}
