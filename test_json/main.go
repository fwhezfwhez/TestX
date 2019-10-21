package main

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/errorx"
)

func main() {
	var js = []byte(`
{"name":"CnsKICAgICJuYW1lIjogImhlbGxvIHdvcmxkIgp9Cg=="}
`)
	type User struct {
		Name string `json:"name"`
	}
	type User2 struct {
		Name []byte `json:"name"`
	}
	var user User
	var user2 User2
    errorx.GroupErrors()
	json.Unmarshal(js, &user)
	json.Unmarshal(js, &user2)
	fmt.Println(string(user.Name))
	fmt.Println(string(user2.Name))
}
