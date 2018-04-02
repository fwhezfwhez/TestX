package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"strings"
)

type Token struct {
 payLoad map[string]string
 header map[string]string
}
func GetToken() *Token{
	token :=Token{}
	token.payLoad = make(map[string]string)
	token.header =make(map[string]string)
	return &token
}
func (token *Token) AddPayLoad(key string,value string) *Token{
	token.payLoad[key]=value
	return token
}

func (token *Token) AddHeader(key string,value string) *Token{
	token.header[key]=value
	return token
}

func (token *Token)JwtGenerator(secretKey string) (jwtResult string,HS256Result string,errorThrow error){
	//1.加密载荷
	payLoad := token.payLoad
	payLoadJson,err := json.Marshal(payLoad)
	if err!=nil {
		fmt.Println(err)
		return "","",err
	}
    payLoadBase64 := base64.StdEncoding.EncodeToString(payLoadJson)
    //fmt.Println("PayLoad的编码结果:",payLoadBase64)

    //2.加密头
	header := token.header
	headerJson,err := json.Marshal(header)
	if err!=nil {
		fmt.Println(err)
		return "","",err
	}
	headerBase64 := base64.StdEncoding.EncodeToString(headerJson)
	//fmt.Println("Header的编码结果:",headerBase64)

	//获得签名
	signature := fmt.Sprintf("%s.%s",payLoadBase64,headerBase64)
	//fmt.Println("签名是:",signature)

	//加密签名
	key:=[]byte(secretKey)
	h:=hmac.New(sha256.New,key)
	h.Write([]byte(signature))
	HS256Rs := base64.StdEncoding.EncodeToString(h.Sum(nil))
	//fmt.Println("经过HS256加密后:",HS256Rs)

	//获得JWT
	jwt :=  fmt.Sprintf("%s.%s",signature,HS256Rs)
	//fmt.Println("jwt是:",jwt)
	return jwt,HS256Rs,nil
}

func (token *Token) Decode(jwt string) (payLoad map[string]string,header map[string]string,HS256Result string,err error){
	jwtArr :=strings.Split(jwt,".")
	payLoadStr := jwtArr[0]
	headerStr := jwtArr[1]
	HS256Rs := jwtArr[2]

	payLoadByte,err := base64.StdEncoding.DecodeString(payLoadStr)
	headerByte,err := base64.StdEncoding.DecodeString(headerStr)
	if err!=nil {
		fmt.Println(err)
		return nil,nil,"",err
	}
	payLoadMap := make(map[string]string)
	headerMap :=make(map[string]string)
	if err:=json.Unmarshal(payLoadByte,&payLoadMap);err!=nil{
		return nil,nil,"",err
	}
	if err:=json.Unmarshal(headerByte,&headerMap);err!=nil{
		return nil,nil,"",err
	}
	return payLoadMap,headerMap,HS256Rs,nil
}

//验证jwt是否合法，解码出的payload和header和服务端的私钥重编码后的HS256值等于解码出的hs值，则合法
//一段jwt合法仅仅是符合私钥，不保证该jwt是否有效，即是否过期!
func(token *Token) IsLegal(jwt string,secretKey string) (bool,error){
	p,h,hs,err :=token.Decode(jwt)
	if err!=nil{
		return false,err
	}
	token.header=h
	token.payLoad =p

	_,HS256,err:=token.JwtGenerator(secretKey)
	if err!=nil{
		return false,err
	}
	if HS256 != hs {
		return false,nil
	}
	return true,nil
}