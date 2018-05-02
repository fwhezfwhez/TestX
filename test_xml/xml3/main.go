package main

import (
	"encoding/xml"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"io/ioutil"
)

//子2节点
type People struct {
	XMLName        xml.Name `xml:"people"`           //
	LinkMan        string   `xml:"link_man"`         //必须真实姓名，不可空
	LinkPhone      string   `xml:"link_phone"`       // 待定
	LinkCreditType string   `xml:"link_credit_type"` //0-身份证，1-学生证，2-军官证，3-护照，4-户口本儿童选择4，5-港澳通行证，6-台胞证,待定
	LinkCreditNo   string   `xml:"link_credit_no"`   //证件号码
}

//子节点
type Peoples struct {
	XMLName xml.Name `xml:"peoples"`
	Pps     []People `xml:"people"`
}
type Conds struct {
	XMLName xml.Name `xml:"conds"`
}

//根标签
type Order struct {
	XMLName       xml.Name `xml:"order"`
	TravelDate    string   `xml:"travel_date"`               //旅游时间，不可空
	EndDate       string   `xml:"end_travel_date,omitempty"` //结束时间，可空
	ArrivedTime   string   `xml:"arrived_time,omitempty"`    //到达时间，可空
	InfoId        string   `xml:"info_id"`                   //产品号，不可空
	CustId        string   `xml:"cust_id"`                   //分销商账号，不可空
	GetType       string   `xml:"get_type,omitempty"`        //配送方式,可空,0免费，1快递费，可空
	OrderSourceId string   `xml:"order_source_id,omitempty"` //对接方系统的订单流水号(唯一)，待定
	OrderMemo     string   `xml:"order_memo,omitempty"`      //订单备注 可空
	Num           string   `xml:"num"`                       //预定数量，不可空

	UserId         string `xml:"user_id,omitempty"`          //用户身份，可空
	LinkMan        string `xml:"link_man"`                   //联系人姓名，用于景区验证，必须真实，不可空
	LinkPhone      string `xml:"link_phone"`                 //接收订单短信或者电话号码 ， 不可空
	LinkEmail      string `xml:"link_email,omitempty"`       //接收订单内容或电子票，待定
	LinkAddress    string `xml:"link_address,omitempty"`     //用于快递，待定
	LinkCode       string `xml:"link_code,omitempty"`        //邮政编码，待定
	LinkCreditType string `xml:"link_credit_type,omitempty"` //联系人证件类型,待定
	LinkCreditNo   string `xml:"link_credit_no,omitempty"`   //联系人证件号码,待定

	Pps Peoples
	Cds Conds
}

//接收调去新增order返回结果
type Result struct{
	XMLName    xml.Name `xml:"result"`
	Status  int `xml:"status"`
	Msg string `xml:"msg"`
	ErrorState string `xml:"error_state"`
	ErrorMsg string `xml:"error_msg"`
	OrderId string `xml:"order_id"`
}

//调去支付返回结果
//<result><orderId></orderId><status><![CDATA[0]]></status><msg>请输入订单号</msg><error_state>10004</error_state><error_msg><![CDATA[支付失败，原因：请=没有输入订单号]]></error_msg></result>
type TextResult struct{
	XMLName    xml.Name `xml:"result"`
	OrderId string `xml:"orderId"`
	Status string  `xml:"status"`
	Msg string	`xml:"msg"`
	ErrorState string	`xml:"error_state"`
	ErrorMsg string		`xml:"error_msg"`
}

func main() {
	order := Order{
		TravelDate: "2018-04-20",
		InfoId:     "22154574",
		CustId:     "1227611",
		OrderMemo:"对接测试",
		Num:        "1",
		LinkMan:    "王杰",
		LinkPhone:  "15507910795",
		OrderSourceId:GetOrderId(),
	}
	formatXml, err := xml.MarshalIndent(order, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println(string(formatXml))

	paramStream, err := xml.Marshal(order)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	v := string(paramStream)

	url:="http://jxyl.ziwoyou.net/api/order.jsp?custId=1227611&apikey=2173031164250D569FF68092358BDB9C&param="+v
	fmt.Println(url)
	fmt.Println(GetDuration(func() {
		DoRequest(url)
	}))

}

var times =0
func DoRequest(url string){
	req,err := http.NewRequest("GET",url,nil)
	if err!=nil {
		fmt.Println(err)
		return
	}
	client:=http.Client{}
	resp,err2:=client.Do(req)
	if err2!=nil{
		fmt.Println(err2)
		return
	}
	rs:=helpRead(resp)
	dest:=Result{}
	err=xml.Unmarshal([]byte(rs),&dest)
	if err!=nil{
		fmt.Println(err)
		return
	}
	if dest.Status==1 {
		doPay(dest)
		return
	}
	if dest.Status!=1&&times!=3{
		times++
		DoRequest(url)
		return
	}
	fmt.Println(times)

}
func doPay(dest Result){
	url:="http://jxyl.ziwoyou.net/api/pay.jsp?custId=1227611&apikey=2173031164250D569FF68092358BDB9C&orderId="+dest.OrderId
	req,err := http.NewRequest("GET",url,nil)
	if err!=nil {
		fmt.Println(err)
		return
	}
	client:=http.Client{}
	resp,err:=client.Do(req)
	if err!=nil {
		fmt.Println(err)
		return
	}
	helpRead(resp)
}
func GetOrderId() string{
	return "D"+time.Now().Format("20060102")+strconv.FormatInt(time.Now().UnixNano()/1e3,10)
}
func GetDuration(f func()) time.Duration {
	t1 := time.Now()
	f()
	t2 := time.Now()
	return t2.Sub(t1)
}
func helpRead(resp *http.Response) string{
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	rs :=string(body)
	fmt.Println(rs)
	return rs
}