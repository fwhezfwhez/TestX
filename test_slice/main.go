package main

import (
	"encoding/json"
	"fmt"
)

type U struct{
	Name string
}
func main(){
	var us []U
	b,_ := json.Marshal(us)
	fmt.Println(string(b))
}
