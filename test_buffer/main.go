package main

import (
	"bytes"
	"fmt"
)

func main() {
	// TestBuffer()
	TestReadPart()
}

// 测试用buffer一次性接受所有数据,不需要预设buffer的初始大小
func TestBuffer() {
	var a = make([]byte, 20)
	var b []byte
	var buffer = bytes.NewBuffer(nil)

	reader := bytes.NewReader(a)

	buffer.ReadFrom(reader)
	b = buffer.Bytes()
	fmt.Println(len(b))

}

func TestReadPart() {
	var a = make([]byte, 20)
	var b = make([]byte, 4, 10)
	reader := bytes.NewReader(a)

	reader.Read(b[0:4])

	buffer := bytes.NewBuffer(nil)
	buffer.ReadFrom(reader)
	fmt.Println(b[3])
	fmt.Println(len(buffer.Bytes()))

	a = make([]byte, 20)
	reader = bytes.NewReader(a)
	buffer.Reset()
	buffer.ReadFrom(reader)


	fmt.Println(len(a))
	fmt.Println(len(buffer.Bytes()))
}
