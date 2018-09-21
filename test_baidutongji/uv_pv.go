package test_baidutongji

import (
	"net/http"
	"fmt"
	"io/ioutil"

	"bytes"
	"encoding/json"
)

func Exec() {
	client := http.Client{}
	type Header struct {
		AccountType int    `json:"account_type"`
		Token       string `json:"token"`
		Username    string `json:"username"`
		Password    string `json:"password"`
	}
	type Body struct {
		SiteId     int    `json:"site_id"`     //站点id // 11842513
		Method     string `json:"method"`      //对应要查询的报告  visit/toppage/a
		StartDate  string `json:"start_date"`  //起始时间 如20160502
		EndDate    string `json:"end_date"`    //结束时间
		Metrics    string `json:"metrics"`     //要查询的条例,多个逗号隔 pv_count,visitor_count,ip_count
		Gran       string `json:"gran"`        //粒度,day/hour/week/month
		MaxResults int    `json:"max_results"` //每页条数，默认20，取0则全部展示
	}

	type Data struct {
		H Header `json:"header"`
		B Body   `json:"body"`
	}
	type RespBody struct{

	}
	type Resp struct{
		Header interface{} `json:"header"`
		RspBody RespBody `json:"body"`
	}


	header := Header{
		Token:"4101703904d81cefc22c71e3f34e4a29",
		Username:"zonstlife",
		Password:"zonst#zonst",
		AccountType:1,
	}
	body := Body{
		SiteId:     11842513,
		Method:     "source/all/a",//"visit/toppage/a",
		StartDate:  "20180823",
		EndDate:    "20180823",
		Metrics:    "pv_count,visitor_count",
		MaxResults: 0,
		Gran:       "day",
	}
	buf, er := json.Marshal(Data{B:body,H:header})
	if er != nil {
		fmt.Println(er.Error())
		return
	}
	req, er := http.NewRequest("POST", "https://api.baidu.com/json/tongji/v1/ReportService/getData", bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/json")
	/*
	    "header": {
        "account_type": 1,
        "password": "你的密码",
        "token": "你的token",
        "username": "你的用户名"
    },
	 */
	if er != nil {
		fmt.Print(er.Error())
		return
	}
	resp, er := client.Do(req)
	if er != nil {
		fmt.Print(er.Error())
		return
	}
	b :=HelpRead(resp)
	var rs interface{}
	er=json.Unmarshal(b,&rs)
	if er != nil {
		fmt.Print(er.Error())
		return
	}
	fmt.Println(rs)
}

func HelpRead(resp *http.Response) []byte{
	buf, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		fmt.Print(er.Error())
		return nil
	}
	rs := string(buf)
	fmt.Println(rs)
	return buf
}
