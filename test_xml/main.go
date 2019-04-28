package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	var a = map[string]interface{}{
		"username": 1,
	}

	b, e := xml.Marshal(a)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))

}
