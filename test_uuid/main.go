package main

import (
	"github.com/satori/go.uuid"
	"fmt"
	"time"
)

func main() {
	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Println(u2.String())
	fmt.Println(u2.String())


	fmt.Println(time.Now().Format("xzczc"))

}
