package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	//"time"
)

func main() {
	//var date = time.Now().Format("2006-01-02")
	//var gameId = 78
	//var secret = "E912JE"
    fmt.Println(fmt.Sprintf("%f", 9.2))
}

func MD5(rawMsg string) string {
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str1)
}
