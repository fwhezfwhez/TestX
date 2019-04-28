package main

import (

	"encoding/binary"
	"fmt"
)

const (
	B  = 1
	KB = 1024 * 1
	MB = 1024 * 1024 * 1
)

func main() {
	var tmp [4]byte

	var a = make([]byte,2)
	fmt.Println(len(append(a, tmp[:]...)))
	binary.BigEndian.PutUint16(a[:] ,uint16(9000))

	rs :=binary.BigEndian.Uint16(a[:])
	fmt.Println(rs)

	fmt.Println(len(a))
	//var data int32 = 728
	//buffer := bytes.NewBuffer(a)
	//e := binary.Write(buffer, binary.BigEndian, data)
	//if e != nil {
	//	panic(e)
	//}
	//var data2 int32
	//e = binary.Read(buffer, binary.BigEndian, &data2)
	//if e != nil {
	//	panic(e)
	//}
	//fmt.Println(data2)

}
