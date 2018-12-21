package main

import (
	"fmt"
	"sync"
)

type U struct{
	Name string
}
var UP = sync.Pool{
	New: func()interface{}{
		return U{"fT"}
	},
}

func main(){
	var u1 = U{"ft"}
	UP.Put(u1)

	u2 := UP.Get().(U)
	u2.Name = "ft2"
	fmt.Println(u2)

	u3 := UP.Get().(U)
	fmt.Println(u3.Name)
}