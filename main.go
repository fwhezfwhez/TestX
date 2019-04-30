package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
}

//公众号支付返回结果结构体
type PayParam struct {
	AppId       string `json:"appId" map:"appId"`
	TimeStamp   string `json:"timeStamp" map:"timeStamp"`
	Status      string `json:"status" map:"status"`
	SignType    string `json:"signType" map:"signType"`
	Package     string `json:"package" map:"package"`
	CallbackURL string `json:"callback_url,omitempty" map:"callback_url"`
	NonceStr    string `json:"nonceStr" map:"status" map:"nonceStr"`
	PaySign     string `json:"paySign" map:"paySign"`
}

func main() {
	handlePayInfo(`{"appId":"wx290ce4878c94369d","timeStamp":"1527220149312","status":"0","signType":"MD5","package":"prepay_id=wx251149092901085a3d4cea760765861054","callback_url":null,"nonceStr":"1527220149312","paySign":"A8189D733CC9BEBADCA6E0C3551DBB11"}`)
}

func handlePayInfo(payInfo string) PayParam {
	payIn := payInfo[1 : len(payInfo)-1]
	payIn = strings.Replace(payIn, "\"", "", -1)
	fmt.Println(payIn)
	params := strings.Split(payIn, ",")
	var temp []string
	var kv = make(map[string]string, 0)
	for _, v := range params {
		temp = strings.Split(v, ":")
		if temp[1] == "" || temp[1] == "nil" || temp[1] == "null" || len(temp) != 2 {
			continue
		}
		kv[temp[0]] = temp[1]
	}
	var payRs = PayParam{
		AppId:     kv["appId"],
		TimeStamp: kv["timeStamp"],
		Status:    kv["status"],
		SignType:  kv["signType"],
		Package:   kv["package"],
		NonceStr:  kv["nonceStr"],
		PaySign:   kv["paySign"],
	}
	fmt.Println(payRs.Package)
	return payRs
}
