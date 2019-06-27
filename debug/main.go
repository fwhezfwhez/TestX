package main

import (
	"fmt"
	"time"
)

func main() {
	t :=time.Now().Add(time.Duration(15) * time.Minute).Add(time.Duration(25) * time.Second).Add(5 * time.Hour)
	fmt.Println(t)
}
