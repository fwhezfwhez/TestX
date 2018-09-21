package main
import (
	"test_X/test_rsa/test_rsa_prv/rsa"
	"crypto"
	"encoding/base64"
	"sort"
	"reflect"
	"strings"
)

//VXWAP wx wap支付rsa加密钥匙
var Public = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApwOQnZ34IZHKQrl8/sgM
vn5TlHNNR/+tAK8EU6OEH7n7NCRsdgzjSxua8GFtAsoXdzDUM8iCXC+/3cxbSQa3
PW8r7sgq5hWZUhQuw0zJFXtXxjLEtmNu1o92pS/C5AGpqNR4Qwa3gNeNUxDYQSII
hy9pmGRFV49ZkrN0/3JRAMlECfPvR2wC6XMR7hOEfA3k5uudDscXUcAaISACVQAb
kbhvx6H+1NKvcb6+m/CdzLNAo6jJYGAePCtN1W9SCYL0FtNc2HfBZ2mK4aik7t6h
4UE/yENJPuRhGQEXWDNylJZBueqnPIvoQ3zWIKLLUsbYAFoOf/11ikkak5nFRHVD
iwIDAQAB
-----END PUBLIC KEY-----
`)

var Private = []byte(`
-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCnA5CdnfghkcpC
uXz+yAy+flOUc01H/60ArwRTo4Qfufs0JGx2DONLG5rwYW0Cyhd3MNQzyIJcL7/d
zFtJBrc9byvuyCrmFZlSFC7DTMkVe1fGMsS2Y27Wj3alL8LkAamo1HhDBreA141T
ENhBIgiHL2mYZEVXj1mSs3T/clEAyUQJ8+9HbALpcxHuE4R8DeTm650OxxdRwBoh
IAJVABuRuG/Hof7U0q9xvr6b8J3Ms0CjqMlgYB48K03Vb1IJgvQW01zYd8FnaYrh
qKTu3qHhQT/IQ0k+5GEZARdYM3KUlkG56qc8i+hDfNYgostSxtgAWg5//XWKSRqT
mcVEdUOLAgMBAAECggEAZaCWtEJ2xb6YDAgg1kxW7Oej5jZPNVuQI1VcVaSvdjrO
cAkl8qOe/4oZuuId2k+s/bs5tRTytjPcFZQPjB1JDVoDqrYk3IMSG87qMa8cwTSf
0Zv89VsGDAuaT6QDfIjCe1hmoaqJqR4BAiXC4cda1x/zqmaBh2lUQR9wNyyRu97B
GL//Jzc4ec7nfFSjkShgsdj3x2xLI/etu8emDSlYue13JbT7rZObKq6K9L0vWuAb
JgriwThwhObjUtJ0pkjWU9QhTTGMa3JCpR6h436u2Wxv6YFzpB3lrAdHiWyXEg7N
ievD6L5yafDKPSutJBPxu55i8ucUUsjwytbGXef/mQKBgQDSP6JSK+sWMCCNruIK
SmJ/cMBXlOir+/uP76XcVXqU67P+kXfrJyE3hkaw+X6H6DuPYZ+RMbeXeVVqcB2R
nVgpsSm1royttXxt8wLJIAEtBduD7SfkX2QqGlwiuEFJLhBT4e4BvzyrWxK+Wd1o
+o7N8MGsb/M4pnQ8hoovwC45HQKBgQDLW3LL2i8CK8czUnhQ9qFm+aJss94orJKA
myugjxOU/oBbFS44qeDWxxeLyYNpMmV2C1X/PCNJYFvnr5FYhqGcMTnYWoo3gKjD
zxngcSuDr2jiMw6f93Ea2JV80I2PUUNKq5aXdBAyvlIsR8NsmrreRuDmS146YyPF
a87PkCn2xwKBgQCO1kSE5ulzzYaQOzu+GgMI6aUC21unXC3v/7QDGN+KxQcn79Q0
ogTjIApuuVHJK3ErW/t4PPqpofjxsvGjt+3phT78+3T/nsjGOc6urVufJteXrUro
0BXQtM/blqpj+TOGiXuVHEMvd1QolgziATd7mhE6AWSQHCXj6NOA1WOR0QKBgQCc
j9NZzqjSaxH6GUw4SSkKNrcuKWyos/Q2v0BfZMgWtbnsZeooBeHseWtPW92oZ/xl
XWTgFFS2cVFQloK4PS6/Kcb3Iv22rX9BAXlSqs5tm+aL//MKhV625Hhyl73k0ROV
sXVHNTwjVRO+s6FpQ9cKulb79pqdKSapH88hTkSAkQKBgQCne3xhicBEL6SkFWXr
yzOmyClTwEoZcEQgVQi5XzuG+Y0OVypzud8/QN9skX+qiQLQtDxm6Poqc3Sigz/P
2IBgNZmmf5JAdtC+r25dNYArYu4jS3jm4/cv33I4p/j6A0qUpXIOaX9jjQHWW71N
cN0K0sSnu//seOrTlGWPD3y/Sg==
-----END PRIVATE KEY-----
`)

func main(){
	rsaClient, err := rsa.NewDefault(string(Private), string(Public))
	handle(err)
	// 签名[]byte格式
	data := ToParam(<微信请求结构体实例>)
	signBytes, err := rsaClient.Sign(data, crypto.SHA256)
	handle(err)
	//签名string 格式
	rs := base64.StdEncoding.EncodeToString(signBytes)



	//解签名与验证
	src := ToParam(<微信返回的结构体实例>)
	signBytesReByte, err := base64.StdEncoding.DecodeString(signReStr)
	handle(err)
	//验证
	errV :=rsaClient.Verify([]byte(src), signBytesReByte, crypto.SHA256, Public)
}


//将该(任意)字符串结构体转换成a=x&b=y&c=z 格式，参数名按照asc递增排序，""值不排
func ToParam(vx interface{}) string {
	var result string
	var tagName_FieldValue = make(map[string]string)
	var SortedArr = make([]string, 0)
	var tagValueTemp string
	var valueStrTemp string

	vType := reflect.TypeOf(vx)
	vValue := reflect.ValueOf(vx)
	for i := 0; i < vType.NumField(); i++ {
		tagValueTemp = filtTag(vType.Field(i).Tag.Get("xml"))
		//设置过滤
		if tagValueTemp == "" || tagValueTemp == "xml" || tagValueTemp == "sign" {
			continue
		}
		valueStrTemp = vValue.Field(i).String()
		if valueStrTemp == "" {
			continue
		}
		tagName_FieldValue[tagValueTemp] = valueStrTemp
		SortedArr = append(SortedArr, tagValueTemp)
	}
	sort.Strings(SortedArr)

	for i, v := range SortedArr {
		if i == 0 {
			result = result + v + "=" + tagName_FieldValue[v]
			continue
		}
		result = result + "&" + v + "=" + tagName_FieldValue[v]
	}
	return result
}

//获取标签第一个值
func filtTag(tag string) string {
	if strings.Contains(tag, ",") {
		return strings.Split(tag, ",")[0]
	} else {
		return tag
	}
}