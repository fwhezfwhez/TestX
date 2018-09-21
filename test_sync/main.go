package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a = make([]int, 0,20)
	a = append(a, 2)

	ToByteArr(&a)
	b := append(a,3)
	ToByteArr(&b)
}

func ToByteArr(a *[]int)[]byte{
	var x reflect.SliceHeader
	var sizeOfArr = int(unsafe.Sizeof(*a))
	fmt.Println("size of a :",sizeOfArr)
	x.Len = sizeOfArr
	x.Cap = sizeOfArr
	x.Data = uintptr(unsafe.Pointer(a))
	fmt.Println("x.Data:",x.Data)
	fmt.Println("x.Len:",x.Data)
	fmt.Println("x.Cap:",x.Data)
	return *(*[]byte)(unsafe.Pointer(&x))
}
func BytesToArr(b []byte) *[]int {
	return (*[]int)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}
