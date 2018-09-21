package main

import (
	"fmt"
	"crypto/md5"
	"unsafe"
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
  var v1 interface{} = (int32)(6)
  var v2 interface{} =(int8)(6)
  fmt.Println(*((*int32)(unsafe.Pointer(&v1))))

  fmt.Println(*((*int64)(unsafe.Pointer(&v1)))==*((*int64)(unsafe.Pointer(&v2))))


}


func MD5(rawMsg string) string{
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	fmt.Println(md5str1)
	return md5str1
}