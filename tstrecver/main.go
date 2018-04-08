package main

import (
	"fmt"
	"crypto/md5"
	"bytes"
	"bufio"
)

const (
	common = "ddddd"
	k1 =common + "dffasdfdasf"
)
type User struct {
	/*
		用户名，用于注册用户和确认用户的唯一性，和一个user对象的name还是有区别的
	 */
	Username string `form:"userName"`
	Password string `form:"password"`
}

func(user *User) EncodingPW(){
	user.Password = MD5(user.Password)
}
func(user User) String() string{
	return fmt.Sprintf("name:%s",user.Username)
}
func main(){
	var user = &User{"ft", "123456"}
	user.EncodingPW()
	fmt.Println(user.Password)

	  user2  :=&User{}
	fmt.Println(user2)

	var userTemp *User
	fmt.Println(userTemp)

   var buf bytes.Buffer
   buf.WriteString("1234")
	buf.WriteString("12354")
   fmt.Println(buf.String())


   fmt.Println(k1)

}


func MD5(rawMsg string) string{
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	fmt.Println(md5str1)
	return md5str1
}