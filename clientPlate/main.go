package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"time"
	"strconv"
	"bytes"
	"encoding/json"
)

//var host = "http://media.qinglong365.com/"
var host = "http://10.0.203.92:8087/"

type Medium struct {
	MediumName  string `json:"mediumName" form:"mediumName"` //媒体称长度在0到20
	Os          string `json:"os" form:"os"`                 //android 或者ios
	Category    string `json:"category" form:"category"`
	SubCategory string `json:"subcategory" form:"subCategory"`
	Keyword     string `json:"keyword,omitempty" form:"keyword"`
	Intro       string `json:"intro,omitempty" form:"intro"`
	UserName    string `json:"userName,omitempty" form:"userName"`
	AppId       string `json:"appId,omitempty" form:"appId"`
}

type Slot struct {
	SlotName   string `json:"slotName" form:"slotName"` //广告位名字,长度在20以内
	UserName   string `json:"userName" form:"userName"`
	MediumName string `json:"mediumName" form:"mediumName"`
	AdType     string `json:"adType" form:"adType"`
	Os         string `json:"os" form:"os"`
	Size       string `json:"size" form:"size"`
	AppId      string `json:"appId,omitempty" form:"appId"`
	SlotId     string `json:"slotId,omitempty" form:"slotId"`
}

var UserName = "admin"
var Category = "游戏"
var SubCategory = "棋牌"
var Intro = "中至软件"
var Keyword = "中至软件"

var AdType = "graphic"
var Size = "300x200"
//var SlotName = "中至广告位"
var Token = ""
//eyJleHAiOiIxNTIyMTE3OTU5Iiwicm9sZSI6ImFkbWluIiwidXNlck5hbWUiOiJhZG1pbiJ9.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.u7J4AZIw1wGrRGezOwkv0o/MGf8DCLbaL6QHBlp/+Uo=
func main() {

	//5.测试用户登录
	//UserLogin()
	//fmt.Println(Token)
	//1.测试用户新增
	//UserInsert()
	//2.测试用户列表
	//UserList()
	//3.测试媒体新增
	//medium := Medium{"testMedium2","ios","消费","娱乐","123","hello","admin","199992"}
	//mediums :=[]Medium{medium}
	//MediumInsert(mediums)
	//4.测试媒体列表
	MediumList()

	//6.获取用户名集
	//GetUserNames()
	//7.获取MediumNames集
	//MediumNames()
	//8.获取medium OS和appId
	//MediumOSAppId()

	//9.实际数据导入
	//PourData()


}

func PourData() {
	var Data = map[string]string{
		"中至南昌麻将ios":     "7",
		"中至南昌麻将android": "7",
		"中至九江麻将ios":     "8",
		"中至九江麻将android": "8",
		"中至上饶麻将ios":     "9",
		"中至上饶麻将android": "9",
		"中至吉安麻将ios":     "11",
		"中至吉安麻将android": "11",
		"中至余干麻将ios":     "28",
		"中至余干麻将android": "28",
		"中至乐平麻将ios":     "29",
		"中至乐平麻将android": "29",
		"中至常熟麻将ios":     "30",
		"中至常熟麻将android": "30",
		"中至德兴麻将ios":     "47",
		"中至德兴麻将android": "47",
		"中至戈阳麻将ios":     "48",
		"中至戈阳麻将android": "48",
		"中至新疆麻将ios":     "55",
		"中至新疆麻将android": "55",
		"中至万年麻将ios":     "59",
		"中至万年麻将android": "59",
	}

	mediums:= Map2MediumSlice(Data)
	slots := GetSlotSlice(mediums)
	MediumInsert(mediums)
	SlotInsert(slots)
}
func Map2MediumSlice(data map[string]string) ([]Medium) {
	mediums := make([]Medium, len(data))
	var i int = 0
	for k, v := range data {
		mediums[i].MediumName = k
		mediums[i].Os = GetOs(k)
		mediums[i].AppId = v
		mediums[i].UserName = UserName
		mediums[i].Category = Category
		mediums[i].SubCategory = SubCategory
		mediums[i].Keyword = Keyword
		mediums[i].Intro = Intro


		i++
	}
	return mediums
}

func GetSlotSlice(mediums []Medium) []Slot{
	slots := make([]Slot, len(mediums))

	for i,v :=range mediums {
		slots[i].UserName = v.UserName
		slots[i].AppId =v.AppId
		slots[i].MediumName = v.MediumName
		slots[i].Os = v.Os
		slots[i].AdType = AdType
		slots[i].Size = Size

		switch  v.MediumName {
		case "中至南昌麻将ios":
			slots[i].SlotId = "10001"
		case "中至南昌麻将android":
			slots[i].SlotId = "10002"
		case "中至九江麻将ios":
			slots[i].SlotId = "10003"
		case "中至九江麻将android":
			slots[i].SlotId = "10004"
		case "中至上饶麻将ios":
			slots[i].SlotId = "10005"
		case "中至上饶麻将android":
			slots[i].SlotId = "10006"
		case "中至吉安麻将ios":
			slots[i].SlotId = "10007"
		case "中至吉安麻将android":
			slots[i].SlotId = "10008"
		case "中至余干麻将ios":
			slots[i].SlotId = "10009"
		case "中至余干麻将android":
			slots[i].SlotId = "10010"
		case "中至乐平麻将ios":
			slots[i].SlotId = "10011"
		case "中至乐平麻将android":
			slots[i].SlotId = "10012"
		case "中至常熟麻将ios":
			slots[i].SlotId = "10013"
		case "中至常熟麻将android":
			slots[i].SlotId = "10014"
		case "中至德兴麻将ios":
			slots[i].SlotId = "10015"
		case "中至德兴麻将android":
			slots[i].SlotId = "10016"
		case "中至戈阳麻将ios":
			slots[i].SlotId = "10017"
		case "中至戈阳麻将android":
			slots[i].SlotId = "10018"
		case "中至新疆麻将ios":
			slots[i].SlotId = "10019"
		case "中至新疆麻将android":
			slots[i].SlotId = "10020"
		case "中至万年麻将ios":
			slots[i].SlotId = "10021"
		case "中至万年麻将android":
			slots[i].SlotId = "10022"
		}

		slots[i].SlotName = v.MediumName + "广告位"
	}
	return slots
}
func GetOs(name string) string {
	var index = 0
	if strings.Contains(name, "ios") {
		index = strings.Index(name, "i")
		return name[index:index+3]
	} else {
		index = strings.Index(name, "a")
		return name[index:index+7]
	}
	return ""
}
func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}
func UserInsert() {
	if true {
		t1 := time.Now()
		var content = fmt.Sprintf("userName=admin&password=123456&name=ft&phone=18985758498&email=145794@qq.com")
		resp, err := http.Post(host+"v1/POST/user/create", "application/x-www-form-urlencoded", strings.NewReader(content))
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}
		helpRead(resp)
	}
}
func UserList() {
	if true {
		t1 := time.Now()
		resp, err := http.Get(host + "v1/GET/users/list")
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}
		helpRead(resp)
	}
}

func MediumInsert(mediums []Medium) {
	//测试媒体新增
	if true {

		//var content  = fmt.Sprintf("mediumName=中至南昌麻将Android&os=Android&category=游戏&subCategory=棋牌&keyword=中至南昌麻将Android&intro=中至南昌麻将Android&userName=admin&appId=7")

		//resp, err := http.Post(host+"v1/POST/medium/create", "application/x-www-form-urlencoded", strings.NewReader(content))
		for i, v := range mediums {
			content, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			t1 := time.Now()
			resp, err := http.Post(host+"v1/POST/medium/create", "application/json", bytes.NewReader(content))
			t2 := time.Now()
			fmt.Println(t2.Sub(t1))
			if err != nil {
				panic(err)
			}
			fmt.Print("第" + strconv.Itoa(i) + "个:")
			helpRead(resp)
		}

	}
}

func MediumList() {
	if true {
		t1 := time.Now()

		//resp, err := http.Get(host + "v1/GET/mediums/list")
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://localhost:8087/v1/GET/mediums/list",nil)
		req.Header.Add("Authorization", "eyJleHAiOiIxNTIyMzg2MTg4Iiwicm9sZSI6ImFkbWluIiwidXNlck5hbWUiOiJhZG1pbiJ9.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.tXfDQ8SVsQFT1YYUU0zOKPcgolK3s7vna6yuKI715Oo=")
		resp, err := client.Do(req)
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}
		helpRead(resp)
	}
}

func SlotInsert(slots []Slot) {
	//测试广告位新增
	if true {
		for i, v := range slots {
			//if strings.Contains(v.MediumName, "常熟") {
			content, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			t1 := time.Now()
			resp, err := http.Post(host+"v1/POST/slot/create", "application/json", bytes.NewReader(content))
			t2 := time.Now()
			fmt.Println(t2.Sub(t1))
			if err != nil {
				panic(err)
			}
			fmt.Print("第"+strconv.Itoa(i)+"个:", v)
			helpRead(resp)
			//}
		}
	}
}

func SlotLsit() {
	//测试广告列表
	if true {
		t1 := time.Now()
		resp, err := http.Get(host + "v1/GET/slots/list?slotName=中间广告")
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}
		helpRead(resp)
	}
}

func UserLogin() {
	if true {

		var content = fmt.Sprintf("userName=admin&password=123456")
		t1 := time.Now()
		resp, err := http.Post(host+"v1/POST/user/login", "application/x-www-form-urlencoded", strings.NewReader(content))
		Token =resp.Header.Get("Authorization")
		fmt.Println(Token)
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}
		helpRead(resp)
	}
}

func GetUserNames() {
	if true {
		t1 := time.Now()
		resp, err := http.Get(host + "v1/GET/users/userNames/list")
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}
		helpRead(resp)
	}
}
func MediumNames() {
	if true {
		t1 := time.Now()
		resp, err := http.Get(host + "v1/GET/mediums/mediumNames/list?userName=admin")
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}

		helpRead(resp)
	}
}

func MediumOSAppId() {

	if true {
		t1 := time.Now()
		resp, err := http.Get(host + "v1/GET/medium/mediumOSAndAppId?mediumName=中至麻将")
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}

		helpRead(resp)
	}
}
