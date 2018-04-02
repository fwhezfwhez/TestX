package main


import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"time"
)

func main(){
	//测试填写用户信息
	//FilInfo()
	//
	//CountVisitor()

	//测试西瓜
	WaterMelon()
}
func FilInfo(){
	t1 := time.Now()
	var content  = fmt.Sprintf("name=fff&phone=15874454847&province=江西&city=南昌&address=aaa&cid=1")
	resp, err := http.Post("http://localhost:8087/travel/user/create", "application/x-www-form-urlencoded", strings.NewReader(content))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	if err!=nil {
		panic(err)
	}
	helpRead(resp)
}
func CountVisitor(){
	t1 := time.Now()

	resp, err := http.Get("http://localhost:8087/travel?cid=1")
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	if err!=nil {
		panic(err)
	}
	helpRead(resp)
}
func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}
func WaterMelon(){
	t1 := time.Now()

	resp, err := http.Get("http://localhost:8087/watermelon?cid=1&actionType=2")
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	if err!=nil {
		panic(err)
	}
	helpRead(resp)
}