package main

import (
	"time"
	"net/http"
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"encoding/xml"
)

var host = "http://localhost:8077"
//var host ="https://market.qinglong365.com"
var maxIn int = 1
var resps []*http.Response = make([]*http.Response, 0)
var errs []error = make([]error, 0)

var client = &http.Client{}

type Order struct {
	Cid       string `json:"cid" form:"cid" superChecker:"cid"`
	Pversion  string `json:"pversion" form:"pversion" superChecker:"pversion"`
	Phone     string `json:"phone" form:"phone" superChecker:"mobilephone|telephone"`
	Uname     string `json:"uname" form:"uname" superChecker:"uname"`
	Money     string `json:"money" form:"money" superChecker:"money"`
	BillType  string `json:"bill_type" form:"bill_type" superChecker:"bill_type"`
	ProductId string `json:"product_id" form:"product_id" superChecker:"product_id"`
	TravelDate string `json:"travel_date" form:"travel_date" superChecker:"travel_date"`
	AppType string `json:"app_type" form:"app_type" superChecker:"app_type"`
	Num string `json:"num" form:"num" superChecker:"num"`
}

type UpdateOrder struct {
	OrderId     string `json:"order_id"`
	OrderStatus string `json:"order_status"`
}

type NotifyResult struct {
	XMLName          xml.Name `xml:"xml"`
	Version          string   `xml:"version"`
	Charset          string   `xml:"charset"`
	SignType         string   `xml:"sign_type"`
	Status           string   `xml:"status"`
	Message          string   `xml:"message"`
	ResultCode       string   `xml:"result_code"`
	MchId            string   `xml:"mch_id"`
	DeviceInfo       string   `xml:"device_info"`
	NonceStr         string   `xml:"nonce_str"`
	ErrCode          string   `xml:"err_code"`
	ErrMsg           string   `xml:"err_msg"`
	Sign             string   `xml:"sign"`
	OpenId           string   `xml:"openid"`
	TradeType        string   `xml:"trade_type"`
	IsSubscribe      string   `xml:"is_subscribe"`
	PayResult        string   `xml:"pay_result"`
	PayInfo          string   `xml:"pay_info"`
	TransactionId    string   `xml:"transaction_id"`
	OutTransactionId string   `xml:"out_transaction_id"`
	SubIsSubscribe   string   `xml:"sub_is_subscribe"`
	SubAppId         string   `xml:"sub_appid"`
	SubOpenId        string   `xml:"sub_openid"`
	OutTradeNo       string   `xml:"out_trade_no"`
	TotalFee         string   `xml:"total_fee"`
	CashFee          string   `xml:"cash_fee"`
	CouponFee        string   `xml:"coupon_fee"`
	FeeType          string   `xml:"fee_type"`
	Attach           string   `xml:"attach"`
	BankType         string   `xml:"bank_type"`
	BankBillNo       string   `xml:"bank_billno"`
	TimeEnd          string   `xml:"time_end"`
}

func main() {
	fmt.Println(GetDuration(func() {
		//测试订单列表
		// TestOrderList()
		//测试错误代码查询
		//TestErrorCode()
		//测试订单插入
		//TestOrderInsert()
		//测试订单更新
		//TestOrderUpdate()
		//测试log
		//TestLog()

		//
		//测试notify
		TestNotify()
	}))
}


func TestNotify(){
	notice := NotifyResult{
		Version:"2.0",
		Charset:"UTF-8",
		SignType:"RSA_1_256",
		Status:"0",
		Message:"",
	}

	//Version          string   `xml:"version"`
	//Charset          string   `xml:"charset"`
	//SignType         string   `xml:"sign_type"`
	//Status           string   `xml:"status"`
	//Message          string   `xml:"message"`
	//ResultCode       string   `xml:"result_code"`
	//MchId            string   `xml:"mch_id"`
	//DeviceInfo       string   `xml:"device_info"`
	//NonceStr         string   `xml:"nonce_str"`
	//ErrCode          string   `xml:"err_code"`
	//ErrMsg           string   `xml:"err_msg"`
	//Sign             string   `xml:"sign"`
	//OpenId           string   `xml:"openid"`
	//TradeType        string   `xml:"trade_type"`
	//IsSubscribe      string   `xml:"is_subscribe"`
	//PayResult        string   `xml:"pay_result"`
	//PayInfo          string   `xml:"pay_info"`
	//TransactionId    string   `xml:"transaction_id"`
	//OutTransactionId string   `xml:"out_transaction_id"`
	//SubIsSubscribe   string   `xml:"sub_is_subscribe"`
	//SubAppId         string   `xml:"sub_appid"`
	//SubOpenId        string   `xml:"sub_openid"`
	//OutTradeNo       string   `xml:"out_trade_no"`
	//TotalFee         string   `xml:"total_fee"`
	//CashFee          string   `xml:"cash_fee"`
	//CouponFee        string   `xml:"coupon_fee"`
	//FeeType          string   `xml:"fee_type"`
	//Attach           string   `xml:"attach"`
	//BankType         string   `xml:"bank_type"`
	//BankBillNo       string   `xml:"bank_billno"`
	//TimeEnd          string   `xml:"time_end"`

	buf,err := xml.Marshal(notice)
	if err!=nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", host+"/v1/POST/VXWAP/notifyUrl", bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var resp *http.Response
	var err2 error
	 GetDuration(func() {
		resp, err2 = client.Do(req)
		helpRead(resp)
	})
}

func TestOrderList() {

	//resp, err := http.Get(host + "v1/GET/mediums/list")

	req, err := http.NewRequest("GET", host+"/v1/GET/order/list?pagesize=20&pageindex=1&orderby=ttime&asc=false&cid=20030&orderid=201804131523591808&uname=石勇&pversion=v1.0&paystatus=1&orderstatus=1", nil)
	//cookie := &http.Cookie{}
	//cookie.Name="Authorization"
	//cookie.Value="eyJleHAiOiIxNTIyNzU4NTk1In0=.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.PF0wTZle+7H5If/BQRLXTp7KEtb/Td3z/S2A7MgNH2Y="
	//req.AddCookie(cookie)
	//req.Header.Add("Authorization", "eyJleHAiOiIxNTIyNzU2MjY5In0=.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.XWdffWT+ceF0tiGUj7ldbg+EoWOx6vlTsoCZnq0sODg=")

	t1 := time.Now()
	var resp *http.Response
	var err2 error
	for i := 0; i < maxIn; i++ {
		resp, err2 = client.Do(req)

		if err2 != nil {
			errs = append(errs, err2)
			continue
		}
		resps = append(resps, resp)
	}
	t2 := time.Now()
	totalTime := t2.Sub(t1)
	aveTime := totalTime / time.Duration(maxIn)
	fmt.Println("总耗时:", totalTime, "平均:", aveTime)
	if err != nil {
		panic(err)
	}
	helpRead(resps[len(resps)-1])
}

func TestErrorCode() {
	t1 := time.Now()
	client := http.Client{}
	req, err := http.NewRequest("GET", host+"/v1/GET/errorCode?errorCode=10030002", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	helpRead(resp)
}

func TestOrderInsert() {
	orderNew := Order{
		Cid : "20030",
		Pversion : "v1.0",
		Phone : "18978376478",
		Uname : "石勇,张三,赵四",
		Money : "10080",
		BillType : "1",
		ProductId :"1002",
		TravelDate :"2018-05-24",
		Num : "3",
		AppType : "1",
	}


	buf, err := json.Marshal(orderNew)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", host+"/v1/POST/order/add", bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var resp *http.Response
	var err2 error
	duration := GetDuration(func() {
		resp, err2 = client.Do(req)
	})
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	helpRead(resp)
	fmt.Println(duration)
}

func TestOrderUpdate() {
	var resp *http.Response
	var er error
	updateOrder := UpdateOrder{}
	updateOrder.OrderId = "D201804261524715488722806"
	updateOrder.OrderStatus = "2"
	uojs, _ := json.Marshal(updateOrder)
	req, err := http.NewRequest("POST", host+"/v1/POST/order/modify", bytes.NewReader(uojs))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}
	duration := GetDuration(func() {
		resp, er = client.Do(req)
	})
	fmt.Println(duration)
	helpRead(resp)
}

func TestLog() {
	req, err := http.NewRequest("GET", host+"/v1/GET/payLog?cid=1001&pversion=v1.0&product_id=1001", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, er := client.Do(req)
	if er != nil {
		fmt.Println(er)
		return
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

func GetDuration(f func()) time.Duration {
	t1 := time.Now()
	f()
	t2 := time.Now()
	return t2.Sub(t1)
}
