package main

import (
	"encoding/json"
	"errorX"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var c http.Client

var count int
var max int = 1772

var lock sync.RWMutex

var wg = sync.WaitGroup{}

func main() {
	randIPTag := genRandomIp()
	userAgent := fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", randIPTag)
	for j := 53; j < 73; j++ {
		for i := 1; i <= 100; i++ {
			time.Sleep(5 * time.Second)
			once(fmt.Sprintf("10.123.%d.%d", j, i), userAgent)
			lock.RLock()
			fmt.Println("已成功票数: ", count)
			lock.RUnlock()

			if count == max {
				fmt.Println("已完成2000票")
				os.Exit(1)
			}
		}
	}
}

func once(ip, userAgent string) {
	// defer wg.Done()
	// gin.Context{}.ClientIP()
	url := "http://www.jxrjjy.com/ajax/vote_h.jsp?cmd=voteItem"
	formData := "vid=4&itemlist=%5B%7B%22itemList%22%3A%5B2%5D%7D%5D&openValidate=false&vCodeId=10984&validateCode=undefined"
	req, e := http.NewRequest("POST", url, strings.NewReader(formData))
	if e != nil {
		fmt.Println(errorx.Wrap(e))
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "http://www.jxrjjy.com/")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Host", "www.jxrjjy.com")
	req.Header.Set("Referer", "http://www.jxrjjy.com/col.jsp?id=122")
	req.Header.Set("Cookie", "_cliid=kIXIJGEySVjM0tYz; _lastEnterDay=2019-07-16; _siteStatId=9efbc773-0528-45e7-9728-08154c8d81de; _siteStatDay=20190716; _siteStatRedirectUv=redirectUv_17221825; _siteStatVisitorType=visitorType_17221825; _siteStatVisit=visit_17221825; _siteStatReVisit=reVisit_17221825; _siteStatVisitTime=1563266664499")
	// Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36
	//req.Header.Set("X-Requested-With", "XMLHttpRequest")

	req.Header.Set("X-Forwarded-For", ip)
	req.Header.Set("X-Real-Ip", ip)
	req.Header.Set("X-Appengine-Remote-Addr", ip)
	req.Header.Set("X-AppEngine-User-IP", ip)

	resp, e := c.Do(req)
	if e != nil {
		fmt.Println(errorx.Wrap(e))
		return
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()

		buf, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			fmt.Println(errorx.Wrap(e))
			return
		}
		fmt.Println(string(buf))
		type Rs struct {
			Success bool `json:"success"`
		}
		var rs Rs
		e = json.Unmarshal(buf, &rs)
		if e != nil {
			fmt.Println(errorx.Wrap(e))
			return
		}
		if rs.Success == true {
			lock.Lock()
			defer lock.Unlock()
			count++
		}
	}
	fmt.Println(resp.Status)
}

func genRandomIp() string {
	var one, two, thr, fou int
	rand.Seed(time.Now().UnixNano())
	one = rand.Intn(200)

	rand.Seed(time.Now().Add(5 * time.Second).UnixNano())
	two = rand.Intn(200)

	rand.Seed(time.Now().Add(5 * time.Second).UnixNano())
	thr = rand.Intn(200)

	rand.Seed(time.Now().Add(5 * time.Second).UnixNano())
	fou = rand.Intn(200)

	return fmt.Sprintf("%d.%d.%d.%d", one, two, thr, fou)
}
