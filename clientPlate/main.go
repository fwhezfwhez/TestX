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

	//var host = "http://media.qinglong365.com/api/"
var host = "http://10.0.205.33:8087/"

type Medium struct {
	MediumName  string `json:"medium_name" form:"medium_name"` //媒体称长度在0到20
	Os          string `json:"os" form:"os"`                 //android 或者ios
	Category    string `json:"category" form:"category"`
	SubCategory string `json:"subcategory" form:"subcategory"`
	Keyword     string `json:"keyword,omitempty" form:"keyword"`
	Intro       string `json:"intro,omitempty" form:"intro"`
	UserName    string `json:"userName,omitempty" form:"userName"`
	AppId       string `json:"appId,omitempty" form:"appId"`
}

type Slot struct {
	SlotName   string `json:"slot_name" form:"slot_name" binding:"required"`    //广告位名字,长度在20以内
	UserName string `json:"username" form:"username" binding:"required"`
	MediumName string `json:"medium_name" form:"medium_name" binding:"required"`
	AdType string `json:"ad_type" form:"ad_type" binding:"required"`
	Os string `json:"os" form:"os" binding:"required"`
	Size string `json:"size" form:"size" binding:"required"`
	AppId string `json:"appid" form:"appid" binding:"required"`
	SlotId string `json:"slot_id" form:"slot_id"`
}

type CidMedium struct{
	ProductId string `json:"product_id" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Category string `json:"category" binding:"required"`
	SubCategory  string `json:"subcategory" binding:"required"`
	Cid string`json:"cid" binding:"required"`
	AppId string `json:"appid"`
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
	//6.获取用户名集
	//GetUserNames()


	//4.测试媒体列表 okok
MediumList()


	//7.获取MediumNames集ok
	//MediumNames()
	//8.获取medium OS和appId OK
	//MediumOSAppId()

	//9.实际数据导入
	//PourData()

	//10.测试cid方式导入
	//InsertByCid()

	//11.测试媒体新增 OK
	//InsertMedium()

	//12.测试广告位查询 OK
	//SlotLsit()

	//13.测试广告位新增
	//InsertSlot2()

	//14.测试cid通道
	//InsertByCid()

	////3.测试媒体新增
	//medium1 := Medium{"渠道1_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "67"}
	//medium2 := Medium{"渠道1_android", "android", "办公", "邮箱", "123", "hello", "admin", "67"}
	//medium3 := Medium{"渠道2_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "68"}
	//medium4 := Medium{"渠道2_android", "android", "办公", "邮箱", "123", "hello", "admin", "68"}
	//medium5 := Medium{"渠道3_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "69"}
	////
	//medium6 := Medium{"渠道3_android", "android", "办公", "邮箱", "123", "hello", "admin", "69"}
	//medium7 := Medium{"渠道4_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "70"}
	//medium8 := Medium{"渠道4_android", "android", "办公", "邮箱", "123", "hello", "admin", "70"}
	//medium9 := Medium{"渠道5_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "71"}
	//medium10 := Medium{"渠道5_android", "android", "办公", "邮箱", "123", "hello", "admin", "71"}
	//
	//
	//medium11 := Medium{"渠道6_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "72"}
	//medium12 := Medium{"渠道6_android", "android", "办公", "邮箱", "123", "hello", "admin", "72"}
	//medium13 := Medium{"渠道7_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "73"}
	//medium14 := Medium{"渠道7_android", "android", "办公", "邮箱", "123", "hello", "admin", "73"}
	//medium15 := Medium{"渠道8_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "74"}
	//
	//medium16 := Medium{"渠道8_android", "android", "办公", "邮箱", "123", "hello", "admin", "74"}
	////medium17 := Medium{"渠道9_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "75"}
	////medium18 := Medium{"渠道9_android", "android", "办公", "邮箱", "123", "hello", "admin", "75"}
	////medium19 := Medium{"渠道10_ios", "ios", "办公", "邮箱", "123", "hello", "admin", "76"}
	////medium20 := Medium{"渠道10_android", "android", "办公", "邮箱", "123", "hello", "admin", "76"}
	//mediums := []Medium{medium15, medium16}
	//MediumInsert(mediums)

	//medium20 := Medium{"test11", "android", "办公", "邮箱", "123", "hello", "ft", "-1"}
	//mediums := []Medium{medium20}
	//MediumInsert(mediums)


	//insertSlot("渠道1_ios","ios","67","11121","广告位1_ios")
	//insertSlot("渠道1_android","android","67","11122","广告位1_android")
	//insertSlot("渠道2_ios","ios","68","11123","广告位2_ios")
	//insertSlot("渠道2_android","android","68","11124","广告位2_android")
	//insertSlot("渠道3_ios","ios","69","11125","广告位3_ios")
	//insertSlot("渠道3_android","android","69","11126","广告位3_android")
	//insertSlot("渠道4_ios","ios","70","11127","广告位4_ios")
	//insertSlot("渠道4_android","android","70","11128","广告位4_android")
	//insertSlot("渠道5_ios","ios","71","11129","广告位5_ios")
	//insertSlot("渠道5_android","android","71","11130","广告位5_android")
	//insertSlot("渠道6_ios","ios","72","11131","广告位6_ios")
	//insertSlot("渠道6_android","android","72","11132","广告位6_android")
	//insertSlot("渠道7_ios","ios","73","11133","广告位7_ios")
	//insertSlot("渠道7_android","android","73","11134","广告位7_android")
	//insertSlot("渠道8_ios","ios","74","11135","广告位8_ios")
	//insertSlot("渠道8_android","android","74","11136","广告位8_android")
	//insertSlot("渠道9_ios","ios","75","11137","广告位9_ios")
	//insertSlot("渠道9_android","android","75","11138","广告位9_android")
	//insertSlot("渠道10_ios","ios","76","11139","广告位10_ios")
	//insertSlot("渠道10_android","android","76","11140","广告位10_android")


}



func InsertByCid(){
	cidM:=CidMedium{
		ProductId:"1007",
		UserName:"ft",
		Category:"办公",
		SubCategory:"邮箱",
		Cid:"1010",
	}
	buf,er := json.Marshal(cidM)
	if er!=nil{
		panic(er.Error())
	}
	req,er:= http.NewRequest("POST",host+"v1/POST/cidChain",bytes.NewReader(buf))
	if er!=nil{
		panic(er.Error())
	}
	req.Header.Set("Content-Type","application/json")

	client :=http.Client{}
	resp,er:=client.Do(req)
	helpRead(resp)
}

func insertSlot(mediumname,os,appId,slotId,slotName string) {
	v := Slot{
		UserName:   "admin",
		AdType:     "graphic",
		Size:       "500x270",
		MediumName: mediumname,
		Os:         os,
		AppId:      appId,
		SlotId:     slotId,
		SlotName:   slotName,
	}
	InsertSlot(v)
}
func InsertSlot2() {
	v := Slot{
		UserName:   "ft",
		AdType:     "graphic",
		Size:       "500x270",
		MediumName: "测试媒体3",
		Os:         "ios",
		AppId:      "-7",
		SlotId:     "-2",
		SlotName:   "测试广告位3",
	}

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
	//fmt.Print("第" + strconv.Itoa(i) + "个:")
	helpRead(resp)
}

func InsertSlot(v Slot) {
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
	//fmt.Print("第" + strconv.Itoa(i) + "个:")
	helpRead(resp)
}
func InsertMedium() {
	v := Medium{
		MediumName:  "测试媒体3",
		Os:          "ios",
		Category:    "办公",
		SubCategory: "邮箱",
		Keyword:     "测试",
		Intro:       "测试",
		UserName:    "ft",
		AppId:       "18",
	}
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
	//fmt.Print("第" + strconv.Itoa(i) + "个:")
	helpRead(resp)
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

	mediums := Map2MediumSlice(Data)
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

func GetSlotSlice(mediums []Medium) []Slot {
	slots := make([]Slot, len(mediums))

	for i, v := range mediums {
		slots[i].UserName = v.UserName
		slots[i].AppId = v.AppId
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
		req, err := http.NewRequest("GET", host+"/v1/GET/mediums/list", nil)
		//req.Header.Add("Authorization", "eyJleHAiOiIxNTIzNDE5MDQ5Iiwicm9sZSI6ImFkbWluIiwidXNlck5hbWUiOiJhZG1pbiJ9.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.wEAxgTXAdJ84zPuCImGxv2pjc6kYB4WelkhBn4+IHHI=")
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
		resp, err := http.Get(host + "v1/GET/slots/list")
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
		Token = resp.Header.Get("Authorization")
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
		resp, err := http.Get(host + "v1/GET/mediums/mediumNames/list?username=ft")
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
		resp, err := http.Get(host + "v1/GET/medium/mediumOSAndAppId?medium_name=中至德兴麻将ios")
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
		if err != nil {
			panic(err)
		}

		helpRead(resp)
	}
}
