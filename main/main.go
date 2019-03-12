package main

import (
	"encoding/binary"
	"fmt"
)

const (
	B = 1
	KB = 1024 * 1
	MB = 1024 * 1024 *1
)
func main() {
	var a  = make([]byte,2)
	binary.BigEndian.PutUint16(a ,uint16(8))
	fmt.Println(a)
	rs :=binary.BigEndian.Uint16(a)
	fmt.Println(rs)
}


func WriteInt16(buf []byte, n int16) {
	buf[0] = byte(n & 0xff)
	buf[1] = byte((n >> 8) & 0xff)
}
