package main

import "fmt"

type U struct{}
func main() {
    var u U
    var u2 = U{}
    fmt.Println(u2 == u)
}

