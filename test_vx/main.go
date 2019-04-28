package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"shangraomajiang/util/common"
	"time"
)

type GetDiamondNumberParam struct {
	Openid      string `json:"openid"`
	Appid       string `json:"appid"`
	Offerid     string `json:"offer_id"`
	Ts          int64  `json:"ts"`
	ZoneId      string `json:"zone_id"`
	Pf          string `json:"pf"`
	Sig         string `json:"sig"`
	AccessToken string `json:"access_token"`
	MpSig       string `json:"mp_sig"`
}
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

type MidasPayResponse struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	BillNo     string `json:"bill_no"`
	Balance    string `json:"balance"`      // 预扣后的余额
	UsedGenAmt int    `json:"used_gen_amt"` // 本次扣的赠送币的金额
}
func main() {
    // officialCase()
    myCase()
}

func officialCase() {
	var accessToken = "19_q4MXN5sD8nRsPd1oICZakPMfgkef-oQhtRCzhBzi5sUHzuX5pVwSU19vsAkRa5PpzzT3BGb-APFyUzfcKBW-NUt5A-6H15EJVJ7qnxTKQA8jqt2D5NNEzOTS74908_xEr7KXAOAQsfSasQwJHCFhAEAAVY"
	param := GetDiamondNumberParam{
		Openid:      "odkx20ENSNa2w5y3g_qOkOvBNM1g",
		Appid:       "wx1234567",
		Offerid:     "12345678",
		Ts:          1507530737,
		ZoneId:      "1",
		Pf:          "android",
		Sig:         "",
		//AccessToken: "tokenvalue",
		MpSig:       "",
	}
	stringA := common.ToParam(param, "json", []string{"sig", "access_token", "mp_sig"}...)
	fmt.Println(stringA == "appid=wx1234567&offer_id=12345678&openid=odkx20ENSNa2w5y3g_qOkOvBNM1g&pf=android&ts=1507530737&zone_id=1")

	stringSignTemp := stringA + "&org_loc=/cgi-bin/midas/sandbox/getbalance&method=POST&secret=zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u"
	sig := getHmacCode(stringSignTemp, "8C0hiJsNb3v6HNPnyZ2Jgaa43t7qn9Oh")
	fmt.Println(sig)

	mg_sig :=getHmacCode(fmt.Sprintf("access_token=%s&appid=wx1234567&offer_id=12345678&openid=odkx20ENSNa2w5y3g_qOkOvBNM1g&pf=android&sig=ca359d9642d69c7535c70bfaffc6452372bf5bd2854fd4981a84692415eda25f&ts=1507530737&zone_id=1"+"&org_loc=/cgi-bin/midas/sandbox/getbalance&method=POST&session_key=V7Q38/i2KXaqrQyl2Yx9Hg==",accessToken), "Lr3MeUAY1Qc0tkX/yU/NPg==", )
	fmt.Println(mg_sig)
}

func myCase() {
	fmt.Println(len("20190328192543-1553772343570435541-33586765-exchange_by_shop-diamond-1"))
	ts := time.Now().Unix()
	fmt.Println("TS:", ts)
	param := MidasPayRequest{
		Openid:      "oZKx35Nm5ztxziIKxmf6jMTmeOpY",
		Appid:       "wxbec7aebf80022eb2",
		OfferId:     "1450019844",
		Ts:          ts,
		ZoneId:      "1",
		Pf:          "android",
		Sig:         "",
		AccessToken: "19_-lHfXfL86j_ClajEJfuV-AZpsDth-WNU34w02Vhv0XVX47dA0k0UcANMFSZ14sfHWCYUnLEKOXSVHTdcMk9Wvn3vR8R68jULkCcLUNreytYPgytuJfwCmrHi78leyERj5kPnkZ0uEk9LJkM_XQPbAFANGF",
		MpSig:       "",
		BillNo: "20190328141036-33586765-exchange_by_shop-diamond-1",
		Amt: 1,
	}
	stringA := common.ToParam(param, "json", []string{"sig", "access_token", "mp_sig","bill_no","amt"}...)
	fmt.Println("stringA:",stringA)

	stringSignTemp := stringA + "&org_loc=/cgi-bin/midas/sandbox/pay&method=POST&secret=8C0hiJsNb3v6HNPnyZ2Jgaa43t7qn9Oh"
	fmt.Println("stringSignTempA:",stringSignTemp)
	fmt.Println("sig:",getHmacCode(stringSignTemp, "8C0hiJsNb3v6HNPnyZ2Jgaa43t7qn9Oh"))
	param.Sig =getHmacCode(stringSignTemp, "8C0hiJsNb3v6HNPnyZ2Jgaa43t7qn9Oh")

	stringB  := common.ToParam(param, "json", []string{"mp_sig","bill_no", "amt"}...)
	stringSignTempB := stringB + "&org_loc=/cgi-bin/midas/sandbox/pay&method=POST&session_key=YeiOaRfiJFmrr/0KVBaBQg=="
	mg_sig :=getHmacCode(stringSignTempB, "YeiOaRfiJFmrr/0KVBaBQg==")
	fmt.Println("stringB:", stringB)
	fmt.Println("stringSignTempB:",stringSignTempB)
	fmt.Println("mg_sig:",mg_sig)
}

func getHmacCode(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//func getSha256Code(s string) string {
//	h := sha256.New()
//	h.Write([]byte(s))
//	return fmt.Sprintf("%x", h.Sum(nil))
//}


