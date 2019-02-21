package main

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"time"
)

//type TimeStamp int64
//type D struct {
//	Name string    `json:"name"`
//	Age  int       `json:"age"`
//	Ts   TimeStamp `json:"ts"`
//}
//
//func (d TimeStamp) MarshalJSON() ([]byte, error) {
//	rs := time.Unix(int64(d), 0).Format("2006-01-02")
//	js, er := json.Marshal(rs)
//	return js, er
//}
//func (d *TimeStamp) UnmarshalJSON(data []byte) error {
//	var rs string
//	e := json.Unmarshal(data, &rs)
//	if e != nil {
//		return e
//	}
//	t, er := time.Parse("2006-01-02", rs)
//	if er != nil {
//		return er
//	}
//	*d = TimeStamp(t.Unix())
//	return nil
//}
//
//type U struct {
//	D2 interface{}
//}
//type U2 struct {
//	Name string `json:"name"`
//}
//type Cla struct {
//	Id int `json:"id"`
//}
//type J struct {
//	U2s  U2
//	Clas Cla
//}
//type J2 struct {
//	Js J
//}
//
//type Us struct{
//	Age int `json:"age"`
//}

type JR struct {
	Name string `json:"name"`
	T    Time   `json:"t"`
}
type Time time.Time
func (t Time) String()string{
	return fmt.Sprintf("%v", time.Time(t))
}
func (t Time) MarshalJSON()([]byte,error){
	return json.Marshal(time.Time(t))
}
func (t *Time) UnmarshalJSON(buf []byte) error {
	var e error
	var tStr string
	var t0 time.Time
	var errors = make([]error, 0, 6)
	e = json.Unmarshal(buf, &t0)
	if e == nil {
		*t = Time(t0)
		return nil
	}
	errors = append(errors, errorx.Wrap(e))
	e = json.Unmarshal(buf, &tStr)
	if e != nil {
		return e
	}

	for _, layout := range []string{
		"2006-01-02",
		"2006-1-2",
		"2006/01/02",
		"2006/1/2",
		"2006.01.02",
		"2006.1.2",
		"2006-01-02 15:04:05",
		"2006-1-2 15:04:05",
	} {
		t0, e = time.Parse(layout, tStr)
		if e == nil {
			*t = Time(t0)
			return nil
		}
		errors = append(errors, errorx.Wrap(e))
	}

	return errorx.GroupErrors(errors...)
}
func main() {
	var t = `{"name":"ft", "t":"2018-01-1"}`
	var jr JR
	e := json.Unmarshal([]byte(t), &jr)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(json.Marshal(jr))
}
