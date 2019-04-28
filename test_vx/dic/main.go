package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shangraomajiang/util/common"
)

func main() {
	var accessToken = "20_1ITnwGsRc2UAwZPunBJTMb7cL20QQrwuFRP-DRPB0NRoqjH3mgkY7eS5JxEe_qjNTQ4FzOK6OKTj7gZOJTNE-AUN-oSJ6BB5KwD6b-wpytLmV30iScb_p9O6soYk0cZ2xUCK2t36tCO_24WAYRKjAIAKDW"
	var openId = "oZKx35Nm5ztxziIKxmf6jMTmeOpY"
	var appId = "wxbec7aebf80022eb2"
	var offerId = "1450019844"
	var offerSecret = "8C0hiJsNb3v6HNPnyZ2Jgaa43t7qn9Oh"
	var sessionKey = "xoKuh9KxeavKBKhw2upz3A=="

	stringTempA := "appid=wxbec7aebf80022eb2&offer_id=1450019844&openid=oZKx35Nm5ztxziIKxmf6jMTmeOpY&pf=android&ts=1553767670&zone_id=1&org_loc=/cgi-bin/midas/sandbox/pay&method=POST&secret=8C0hiJsNb3v6HNPnyZ2Jgaa43t7qn9Oh"
	stringTempB := "access_token=20_1ITnwGsRc2UAwZPunBJTMb7cL20QQrwuFRP-DRPB0NRoqjH3mgkY7eS5JxEe_qjNTQ4FzOK6OKTj7gZOJTNE-AUN-oSJ6BB5KwD6b-wpytLmV30iScb_p9O6soYk0cZ2xUCK2t36tCO_24WAYRKjAIAKDW&appid=wxbec7aebf80022eb2&offer_id=1450019844&openid=oZKx35Nm5ztxziIKxmf6jMTmeOpY&pf=android&sig=efb5e80dbda1d97ae7cd12f03b4510593c5efd922a744231a7e0d9a5bb4cfca6&ts=1553767670&zone_id=1&org_loc=/cgi-bin/midas/sandbox/pay&method=POST&session_key=xoKuh9KxeavKBKhw2upz3A=="
	type MidasPayRequest struct {
		Openid      string `json:"openid"`
		Appid       string `json:"appid"`
		OfferId     string `json:"offer_id"`
		Ts          int64  `json:"ts"`
		ZoneId      string `json:"zone_id"`
		Pf          string `json:"pf"`
		Amt         int    `json:"amt"`     // 不能为0
		BillNo      string `json:"bill_no"` // 唯一
		Sig         string `json:"sig"`
		AccessToken string `json:"access_token,omitempty"` //请求时消除
		MpSig       string `json:"mp_sig"`
	}
	param := MidasPayRequest{
		Openid:      openId,
		Appid:       appId,
		OfferId:     offerId,
		Ts:          1553767670,
		ZoneId:      "1",
		Pf:          "android",
		Amt:         1,
		BillNo:      "ft",
		AccessToken: accessToken,
	}

	param.Sig = common.HmacHs256(stringTempA, offerSecret)
	param.MpSig = common.HmacHs256(stringTempB, sessionKey)

	buf, _ := json.Marshal(param)

	req, e := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/midas/sandbox/pay?access_token="+accessToken, bytes.NewReader(buf))
	if e != nil {
		panic(e)
	}
	var c = http.Client{}
	resp, e := c.Do(req)
	rs, e := ioutil.ReadAll(resp.Body)
	fmt.Println(string(rs))

}
