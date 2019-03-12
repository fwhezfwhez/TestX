package main

import "fmt"

type Connection int
type Handler func(conn *Connection) bool

func main() {
	var a = func() func(conn *Connection) bool {
		var c = make(map[string]interface{}, 0)
		return func(conn *Connection) bool {
            c["1"] = "ft"
            fmt.Println(c["1"])
            return true
		}
	}

	var con = Connection(1)
	a()(&con)
}
