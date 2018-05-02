package main

import(
	"fmt"
	"time"
	"encoding/xml"
)

type TNote struct {
	Lang    string `xml:"lang,attr"`
	Content string `xml:",innerxml"`
}

type TFile struct {
	XMLName  struct{} `xml:"file"`
	FileName string   `xml:"name,attr"`
	Size     string   `xml:"size,attr"`
}

type Release struct {
	XMLName   struct{} `xml:"release"`
	Version   string   `xml:"version,attr"`
	TimeStamp string   `xml:",attr"`
	Lang      string   `xml:"-"`
	Skin      string   `xml:",chardata"`
	Site      string   `xml:",omitempty"`
	File      []TFile  `xml:",innerxml"`
	CnNotes   TNote    `xml:"cnnote"`
	EnNotes   TNote    `xml:"ennote"`
	Comment   string   `xml:",comment"`
	DDate time.Time `xml:"date,omitempty"`
}

func main(){
	release := Release{Version:   "1.0.0.0",
		TimeStamp: time.Now().String(),
		Lang:      "zh-cn",
		Site:      "",
		File: []TFile{TFile{FileName: "/deploy/package_1.zip", Size: "50"},
			TFile{FileName: "/deploy/package_2.zip", Size: "60"},
		},
		Skin:    "blue",
		Comment: "this is a test for xml parser.",
	}

	v, err := xml.MarshalIndent(release, "", "         ")
	if err != nil {
		fmt.Println("marshal xml value error, error msg:%s", err.Error())
	}

	fmt.Println( string(v))

	v2,err2 :=xml.Marshal(release)
	if err2 != nil {
		fmt.Println("marshal xml value error, error msg:%s", err.Error())
		return
	}
	fmt.Println(string(v2))
}