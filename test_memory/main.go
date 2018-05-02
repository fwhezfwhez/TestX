package main

import (
	"unsafe"
	"fmt"
)

type User struct{
	Id int
	Name string
	Married bool
	Salary float64
	Error error
}

func (user User) Sizeof(){
	fmt.Println("Id,int,",unsafe.Sizeof(user.Id))
	fmt.Println("Name,string,",unsafe.Sizeof(user.Name))
	fmt.Println("Married,bool",unsafe.Sizeof(user.Married))
	fmt.Println("Salary,float64,",unsafe.Sizeof(user.Salary))
	fmt.Println("Error,error,",unsafe.Sizeof(user.Error))
}
func main() {
	user := User{}
	user.Sizeof()
	fmt.Println(unsafe.Sizeof(user))
}