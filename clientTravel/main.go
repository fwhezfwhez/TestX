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
	//WaterMelon()

	//获取令牌
	//Login()
	//测试下载
	DownLoad()
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

func Login(){
	t1 := time.Now()
	var content  = fmt.Sprintf("userName=邹梦君&password=123456")
	resp, err := http.Post("http://localhost:8087/login", "application/x-www-form-urlencoded", strings.NewReader(content))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	if err!=nil {
		panic(err)
	}
	helpRead(resp)
}
func DownLoad(){
	t1 := time.Now()

	//resp, err := http.Get(host + "v1/GET/mediums/list")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8087/download",nil)
	cookie := &http.Cookie{}
	cookie.Name="Authorization"
	cookie.Value="eyJleHAiOiIxNTIyNzU4NTk1In0=.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.PF0wTZle+7H5If/BQRLXTp7KEtb/Td3z/S2A7MgNH2Y="
	req.AddCookie(cookie)
	//req.Header.Add("Authorization", "eyJleHAiOiIxNTIyNzU2MjY5In0=.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.XWdffWT+ceF0tiGUj7ldbg+EoWOx6vlTsoCZnq0sODg=")
	resp, err := client.Do(req)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	if err != nil {
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