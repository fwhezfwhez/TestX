package main

import (
	"encoding/xml"
	"fmt"
)

type Message struct {
	Xml          xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId"`
}

func main(){
	str := `<xml> 
				<ToUserName><![CDATA[toUser]]></ToUserName>  
				<FromUserName><![CDATA[fromUser]]></FromUserName> 
				<CreateTime>1348831860</CreateTime>  
				<MsgType><![CDATA[text]]></MsgType> 
			    <Content><![CDATA[this is a test]]></Content> 
				<MsgId>1234567890123456</MsgId>  
			</xml>`
	obj := Message{}
	err := xml.Unmarshal([]byte(str),&obj)
	if err!=nil{
		panic(err)
	}
	fmt.Println(obj)
	rs,_:=xml.MarshalIndent(obj,"","    ")
	fmt.Println(string(rs))
}