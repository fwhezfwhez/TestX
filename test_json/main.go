package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeStamp int64
type D struct {
	Name string    `json:"name"`
	Age  int       `json:"age"`
	Ts   TimeStamp `json:"ts"`
}

func (d TimeStamp) MarshalJSON() ([]byte, error) {
	rs := time.Unix(int64(d), 0).Format("2006-01-02")
	js, er := json.Marshal(rs)
	return js, er
}
func (d *TimeStamp) UnmarshalJSON(data []byte) error {
	var rs string
	e := json.Unmarshal(data, &rs)
	if e != nil {
		return e
	}
	t, er := time.Parse("2006-01-02", rs)
	if er != nil {
		return er
	}
	*d = TimeStamp(t.Unix())
	return nil
}

type U struct {
	D2 interface{}
}
type U2 struct {
	Name string `json:"name"`
}
type Cla struct {
	Id int `json:"id"`
}
type J struct {
	U2s  U2
	Clas Cla
}
type J2 struct {
	Js J
}

type Us struct{
	Age int `json:"age"`
}
func main() {
	var u Us
	e:=json.Unmarshal([]byte(`
		{
			"name":"ft",
			"age":9
		}
		`),&u)
	if e!=nil {
		panic (e)
	}
	fmt.Println(u)
}


