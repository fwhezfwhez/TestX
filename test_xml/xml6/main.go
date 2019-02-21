package main

import (
	"encoding/xml"
	"fmt"
)

type ContactMechLst struct {
	XMLName     xml.Name `xml:"ns1:ContactMechLst"`
	Ns1         string   `xml:"xmlns:ns1,attr"`
	ContactMech []ContactMech
}

type ContactMech struct {
	XMLName xml.Name  `xml:"ns1:ContactMech"`
	Telcom  Telcom
	Email   Email
	Address []Address
}

type Address struct {
	XMLName  xml.Name `xml:"ns1:Address"`
	AddType  string   `xml:"ns1:AddType"`
	AddrL1   string   `xml:"ns1:AddrL1"`
	AddrL2   string   `xml:"ns1:AddrL2"`
	AddrL3   string   `xml:"ns1:AddrL3"`
	City     string   `xml:"ns1:city"`
	PostCode string   `xml:"ns1:PostCode"`
	Country  string   `xml:"ns1:Country,omitempty"`
}
type Email struct {
	XMLName xml.Name `xml:"ns1:Email"`

	EmailType string `xml:"ns1:EmailType,omitempty"`
	EAddr     string `xml:"ns1:EAddr,omitempty"`
}

type Telcom struct {
	XMLName       xml.Name `xml:"ns1:Telcom"`
	ContactNumber string   `xml:"ns1:ContactNumber,omitempty"`
}

func main() {
	var c = ContactMechLst{
		Ns1: "http://www.fisglobal.com/services/prepaid/common",
		ContactMech: []ContactMech{
			{
				Telcom: Telcom{
					ContactNumber: "+919084258647",
				},
			},
			{
				Email: Email{
					EmailType: "1",
					EAddr:     "priya@fisglobal.com",
				},
			},
			{
				Address: []Address{
					{
						AddType:  "1",
						AddrL1:   "RK Road",
						AddrL2:   "Mylapore",
						AddrL3:   "Chennai",
						City:     "TamilNadu",
						PostCode: "600200",
						Country:  "356",
					},
					{
						AddType:  "2",
						AddrL1:   "No.22, Sardar Patel Road",
						AddrL2:   "Guindy",
						AddrL3:   "Chennai",
						City:     "TamilNadu",
						PostCode: "600032",
					},
				},
			},
		},
	}
	buf, e := xml.MarshalIndent(c, "", "    ")
	if e != nil {
		panic(e)
	}
	fmt.Println(string(buf))
}
