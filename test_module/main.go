package main

import (
	"encoding/json"
	"fmt"
)

type Link struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Target string `json:"target"`
	Icon   string `json:"icon"`
}

type SubItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Link Link   `json:"link"`
}
type MenuItem struct {
	Name string    `json:"name"`
	Link Link      `json:"link"`
	Sub  []SubItem `json:"sub"`
}

type Menu struct {
	Menu []MenuItem `json:"menu"`
}

func main() {
	data := `{
	"menu": [
            {
                "name": "首页",
                "link": {
                    "type": "url",
                    "value": "http://www.baidu.com",
                    "target": "",
                    "icon":"1.jpg"
                }
            },
            {
                "name": "关于",
                "link":{
                    "type":"empty"
                },
                "sub": [
                    {
                        "id": "page3",
                        "name": "公司",
                        "link": {
                            "type": "page",
                            "value": "111",
                            "target": "_black"
                        }
                    },
                    {
                        "id": "page3",
                        "name": "公司",
                        "link": {
                            "type": "apply",
                            "value": "apply_id/category_id",
                            "target": "_black"
                        }
                    }
                ]
            }
        ]
	}`
	menu := Menu{}
	er := json.Unmarshal([]byte(data), &menu)
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	fmt.Println(menu)
	fmt.Println(menu.Menu)

	x,er :=json.Marshal(menu.Menu)
	if er!=nil{
		fmt.Println(er.Error())
		return
	}
	fmt.Println(string(x))
}
