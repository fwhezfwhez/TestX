package main

import (
	"net/http"
	//"io/ioutil"
	"fmt"
	//"github.com/gin-gonic/gin/json"
	//"bytes"

	"io/ioutil"
	"encoding/xml"
	"bytes"
	"encoding/json"
)

type Xml struct {
	XMLName          xml.Name `xml:"xml"`
	Age string
	Year int
}
type User struct{
	Name string `json:"uname"`
	Age string `json:"id"`
}
func main() {
   //buf:=[]byte("hello")
   //fmt.Println("buf:",string(buf))
   //reader:= bytes.NewReader(buf)
   ////var result = make([]byte,len(buf))
   ////reader.Read(result)
   //result,_ := ioutil.ReadAll(reader)
   //fmt.Println("rs：",string(result))

	//defer func(func ()){}(func (){fmt.Println("ok")})
	//fmt.Println(5)

	user := User{"ft","9"}
	usJS,er:=json.Marshal(user)
	if er!=nil {
		fmt.Println(er.Error())
		return
	}
	req,err:=http.NewRequest("DELETE","http://localhost:8088/user/{fttt}/{5}/delete",bytes.NewReader(usJS))
	req.Header.Set("Content-Type", "application/json")
	if err!=nil{
		fmt.Println(err)
		return
	}
	client :=http.Client{}
	resp,err:=client.Do(req)
	if err!=nil{
		fmt.Println(err)
		return
	}
	helpRead(resp)


	//req,err:=http.NewRequest("GET","http://172.21.166.81:8088/TP",nil)
	////req.Header.Set("Content-Type", "application/json")
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//client :=http.Client{}
	//resp,err:=client.Do(req)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//helpRead(resp)

	//xxml:=Xml{
	//	Age:"33",
	//	Year:1000,
	//}
	//buf,err:=xml.Marshal(xxml)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//req,err:=http.NewRequest("POST","http://10.0.203.92:8087/Test",bytes.NewReader(buf))
	//req.Header.Set("Content-Type", "application/xml")
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//client :=http.Client{}
	//resp,err:=client.Do(req)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//helpRead(resp)
	//exp, errr := time.Parse("2006-01-02 15:04:05", time.Now().Add(2*time.Hour).Format("2006-01-02 15:04:05"))
	//if errr!=nil{
	//	fmt.Println(errr)
	//}
	//fmt.Println(exp)
	//url := "http://10.0.203.92:8087/Test"
	//req, _ := http.NewRequest("GET", url, nil)
	////req.Header.Add("x-auth-token", "abc")
	////req.Header.Add("x-auth-token","abc2")
	//res, _ := http.DefaultClient.Do(req)
	//defer res.Body.Close()
	////body, _ := ioutil.ReadAll(res.Body)
	//helpRead(res)
	//
	//fmt.Println(res.Header.Get("x-auth-token"))
	////fmt.Println(string(body))
	//
	//type User struct{
	//	Name *string
	//}


}

func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}