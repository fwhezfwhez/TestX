package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//fmt.Println(os.Getenv("GORM_DSN"))
	var json1 = `
		{
			"type":"football",
			"receive":{
                 "name":"足球",
                 "start_at":"2018-01-01"
             }
		}
    `
	var json2 = `
		{
			"type":"basketball",
			"receive":{
                 "game_name":"篮球球",
                 "start_at":"2018-01-01",
                 "fee":"3000"
             }
		}
    `
	type R struct {
		VType   string          `json:"type"`
		Receive json.RawMessage `json:"receive"`
	}
	type FootBall struct {
		Name    string `json:"name"`
		StartAt string `json:"start_at"`
	}
	type BasketBall struct {
		GameName string `json:"game_name"`
		StartAt  string `json:"start_at"`
		Fee      int    `json:"fee,string"`
	}

	var r = R{}

	// football json
	{
		e := json.Unmarshal([]byte(json1), &r)
		if e != nil {
			panic(e)
		}
		var footBall = FootBall{}
		e = json.Unmarshal(r.Receive, &footBall)
		if e != nil {
			panic(e)
		}
		fmt.Println(footBall)
	}
	// basketball json
	{
		e := json.Unmarshal([]byte(json2), &r)
		if e != nil {
			panic(e)
		}
		var basketBall = BasketBall{}
		e = json.Unmarshal(r.Receive, &basketBall)
		if e != nil {
			panic(e)
		}
		fmt.Println(basketBall)
	}

}
