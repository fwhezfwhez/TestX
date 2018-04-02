package main

import (
	"net/http"
	//"io/ioutil"
	"fmt"
	"time"
)

func main() {
	exp, errr := time.Parse("2006-01-02 15:04:05", time.Now().Add(2*time.Hour).Format("2006-01-02 15:04:05"))
	if errr!=nil{
		fmt.Println(errr)
	}
	fmt.Println(exp)
	url := "http://10.0.203.92:8087/Generate"
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Add("x-auth-token", "abc")
	//req.Header.Add("x-auth-token","abc2")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.Header.Get("x-auth-token"))
	//fmt.Println(string(body))
}