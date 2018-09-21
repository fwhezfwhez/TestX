package main

import (
	"net/url"
	"fmt"
)

func main() {
	fmt.Println(url.QueryEscape("3123"))
	urlRaw := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx3210548e5292e030&redirect_uri=https%3A%2F%2Fmarket.qinglong365.com%2Fv1%2FGET%2FVXPub%2FGetOpenId%2F&response_type=code&scope=snsapi_base&state=D201805211526868819671342#wechat_redirect"
	fmt.Println(len("9d101c97133837e13dde2d32a5054abb")==len("9df7aec7bbc9c90c08608eff33184f21"))

	fmt.Println("https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx3210548e5292e030&redirect_uri=https%3A%2F%2Fmarket.qinglong365.com%2Fv1%2FGET%2FVXPub%2FGetOpenId%2F&response_type=code&scope=snsapi_base&state=D201805211526868819671342#wechat_redirect"==urlRaw)

	var m = make(map[string][]string,0)
	m["1"]=[]string{"1","2","3"}

	t:= m["1"]
	t= append(t,"4")
	m["1"]=t
	fmt.Println(m)
	}