package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"crypto/md5"
	"encoding/json"
	"strings"
)

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

type KeySecret struct {
	AppKey    string
	AppSecret string
}
//url =
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


func main() {

	//1.测试缓存User的appkey和appsecret
	//TestUserKeySecret()

	//2.测试medium缓存
	//TestMediumRedis()

	////3.测试slot缓存
	TestSlotRedis()

	//4.测试清理缓存
	// TestDel()
}

func MD5(rawMsg string) string {
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}
func TestUserKeySecret(){

		c, err := redis.DialURL("redis://localhost:6379")
		if err != nil {
			fmt.Println("1.连接redis失败", err)
			return
		}
		fmt.Println("1.连接成功")
		defer c.Close()
		userName := "admin"
		key := MD5(userName)
		result, err := redis.Bytes(c.Do("GET", key))
		if err != nil {
			panic(err)
		}
		ks := KeySecret{}
		err = json.Unmarshal(result, &ks)
		if err != nil {
			panic(err)
		}
		fmt.Println(ks)

}

func  TestMediumRedis(){
	if true{
		mediums:= Map2MediumSlice(Data)
		c, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println("1.连接redis失败", err)
			return
		}
		fmt.Println("1.连接成功")
		defer c.Close()
		for i:=0;i<len(mediums);i++ {
			result, err := redis.Bytes(c.Do("HGET", "app_list", "1gXD51HotduqioLj:"+mediums[i].Os+":"+mediums[i].AppId))
			fmt.Println("bk9LRewrTax8yYtn:"+mediums[i].Os+":"+mediums[i].AppId)
			if err != nil {
				fmt.Println(err)
			}
			var target interface{}
			json.Unmarshal(result, &target)
			fmt.Println(target)
		}
	}
}

func TestDel(){

	mediums:= Map2MediumSlice(Data)
	slots := GetSlotSlice(mediums)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("1.连接redis失败", err)
		return
	}
	fmt.Println("1.连接成功")
	defer c.Close()
	for i:=0;i<len(mediums);i++{
		_,err=c.Do("HDEL","app_list","UUt7dXUB8bJ2r5yC:"+mediums[i].Os+":"+mediums[i].AppId)
		if err!=nil {
			fmt.Println(err)
		}
		_,err = c.Do("HDEL","app_slot","UUt7dXUB8bJ2r5yC:"+slots[i].Os+":"+slots[i].AppId+":"+slots[i].SlotId)
		if err!=nil {
			fmt.Println(err)
		}
	}
	fmt.Println("删除缓存成功")

}

func TestSlotRedis(){
	if true{
		c, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println("1.连接redis失败", err)
			return
		}
		fmt.Println("1.连接成功")
		defer c.Close()

		mediums:= Map2MediumSlice(Data)
		slots := GetSlotSlice(mediums)
		for i:=0;i<len(slots);i++ {
			result, err := redis.Bytes(c.Do("HGET", "app_slot", "1gXD51HotduqioLj:"+slots[i].Os+":"+slots[i].AppId+":"+slots[i].SlotId))
			fmt.Println("1gXD51HotduqioLj:"+slots[i].Os+":"+slots[i].AppId+":"+slots[i].SlotId)
			if err != nil {
				fmt.Println(err)
			}
			var target interface{}
			json.Unmarshal(result, &target)
			fmt.Println(target)
		}
	}
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