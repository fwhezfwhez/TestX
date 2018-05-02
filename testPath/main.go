package main

import (
"fmt"
	"runtime"
	"path"
)

type User struct{
	Count int
}
func main() {
	user :=User{1}
	Add(&user)
	fmt.Println(user.Count)
	str1:=Pa()
	fmt.Println(str1)
}
func Pa() string{
	_,file,line,_:=runtime.Caller(1)
	fmt.Println(file,line)
	dir :=path.Dir(file)
	return dir
}
func Add(u *User){
	u.Count++
}
