package main

import (
	"fmt"
	"go/types"
	"container/list"
	"time"
)
type User struct{}
func main() {
  var  users = make([]User,0)
   var user = User{}
   te(user)
   te(users)

   m:=make(map[string]string,5000)
  // m["d"]="f"
   delete(m,"d")
   fmt.Print(m)

   var l = &list.List{}
   e:=l.PushFront(5)
	l.PushFront(6)
   l.Remove(e)
	l.Remove(e)
  fmt.Println(l.Len())

  fmt.Println(time.Now().Day())

  var c = make([]string,0,10)
  c[1]="5"

}

func te(dest interface{}){
	switch dest.(type){
	case types.Pointer:
		fmt.Println("是切片")
	case types.Struct:
		fmt.Println("是实例")
	default:
		fmt.Println("都不是")
	}
}