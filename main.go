package main

import (
	"fmt"
	"go/types"
	"container/list"
	"time"
	"runtime"

)
type User struct{
	Name string
}
func main() {
  var  users = make(map[string]User,0)
   var user = User{"ft"}
   //te(user)
   //te(users)
   users["1"] = user
   fmt.Println(users["1"].Name)

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


  fmt.Println(runtime.NumCPU())
  pre := runtime.GOMAXPROCS(runtime.NumCPU())
  fmt.Println(pre,runtime.NumCPU())

  ms:=make(map[int]*User,0)
  ms[1]= &user
  fmt.Println(len(ms))
  ms[1]=nil
  fmt.Println(len(ms))
  fmt.Println(ms[1])

  var i=5
  switch i {
  case 5:
  	fmt.Println(1)
  	fmt.Println(2)
  case 3:
  	fmt.Println(3)
  }
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