package main


import (
	"encoding/xml"
	"fmt"

)

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}
type Result struct{
	XMLName    xml.Name `xml:"result"`
	Status  int `xml:"status"`
	Msg string `xml:"msg"`
	ErrorState string `xml:"error_state"`
	ErrorMsg string `xml:"error_msg"`
	OrderId string `xml:"order_id"`
}

func main() {
	//str:="<result><status><![CDATA[0]]></status><msg><![CDATA[2018-04-25获取价格失败，已经销售完毕或者取消班期。]]></msg><error_state>10001</error_state><error_msg><![CDATA[产品已经无法获取当前日期的价格]]></error_msg></result>"
	str:="<result><status><![CDATA[1]]></status><msg><![CDATA[录入订单成功！本订单需要在线支付后才能有效！]]></msg><error_state>10000</error_state><error_msg><![CDATA[录入订单成功！本订单需要在线支付后才能有效！]]></error_msg><order_id><![CDATA[51252546]]></order_id><order_money><![CDATA[32]]></order_money><mem_order_money><![CDATA[39.9]]></mem_order_money><order_state><![CDATA[1]]></order_state></result>"
	data := []byte(str)
	v := Result{}
	err := xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	rs := fmt.Sprintf("state:%d,msg:%s,errCode:%s,errMsg:%s,orderId:%s",v.Status,v.Msg,v.ErrorState,v.ErrorMsg,v.OrderId)
	fmt.Println(rs)
}