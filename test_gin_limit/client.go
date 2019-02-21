package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{}
	buf := `?start_time=2010-01-01&end_time=2010-01-03&time_format=2006-01-02&state=1`
	req, e := http.NewRequest("GET",  fmt.Sprintf("http://localhost:8892/?%s", buf),nil)
	if e != nil {
		panic(e)
	}
	resp, e := c.Do(req)
	if e != nil {
		panic(e)
	}
	rs, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(rs))

	m1,_:=json.Marshal(time.Now())
	fmt.Println(string(m1))
}
