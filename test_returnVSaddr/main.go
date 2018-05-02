package main
type User struct{
	Name string
}
func Addr(user *User){
	user.Name ="changed"
}

func Return(user User) User{
	user.Name= "changed"
	return user
}
