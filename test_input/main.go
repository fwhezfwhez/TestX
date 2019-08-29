package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f := bufio.NewReader(os.Stdin)
	for {
		input, _ := f.ReadString('\n')
		if len(input) == 1 {
			continue
		}
		fmt.Println(input)
	}
}
