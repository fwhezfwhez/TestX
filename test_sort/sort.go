package main

import (
	"fmt"
	"sort"
)

type User struct{
	Id int
	Salary float32
	Name string
}
type  Users []User

func (users Users) Less(i int ,j int) bool{
	return  users[i].Salary<users[j].Salary
}
func (users Users) Swap(i,j int){
	users[i],users[j] = users[j], users[i]
}
func  (users Users) Len() int {
	return len(users)
}
func main(){
	var users = Users{
		{1,3000,"ft"},
		{Id:2,Salary:4000,Name:"ft2"},
		{3,2000,"ft3"},
	}
	fmt.Println("原始排序:",users)
	sort.Sort(users)
	fmt.Println("排序后:",users)
	//reverse返回的是sort接口对象，所以还要包一层Sort方法
	sort.Sort(sort.Reverse(users))
	fmt.Println("倒置后:",users)

	var str = []string{
		"ac","cd","bd",
	}
	sort.Strings(str)
	fmt.Println(str)

}