package main

import (
	"fmt"
	"crypto"
	"github.com/pkg/errors"
	"unicode/utf8"
	"github.com/grpc/grpc-go/grpclog"
)

var private = `
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
`
var public = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApwOQnZ34IZHKQrl8/sgM
vn5TlHNNR/+tAK8EU6OEH7n7NCRsdgzjSxua8GFtAsoXdzDUM8iCXC+/3cxbSQa3
PW8r7sgq5hWZUhQuw0zJFXtXxjLEtmNu1o92pS/C5AGpqNR4Qwa3gNeNUxDYQSII
hy9pmGRFV49ZkrN0/3JRAMlECfPvR2wC6XMR7hOEfA3k5uudDscXUcAaISACVQAb
kbhvx6H+1NKvcb6+m/CdzLNAo6jJYGAePCtN1W9SCYL0FtNc2HfBZ2mK4aik7t6h
4UE/yENJPuRhGQEXWDNylJZBueqnPIvoQ3zWIKLLUsbYAFoOf/11ikkak5nFRHVD
iwIDAQAB
-----END PUBLIC KEY-----
`
func main() {
	fmt.Println("hello rsa")
	data := "hello rsa"

	rs,err:=RsaSign(data,"RSA_1_256",private)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rs)
}

func RsaSign(data string,signType string,privateKey string) (string,error){
	//加密方式判断
	var suit crypto.Hash
	if signType == "RSA_1_256"{
		suit =crypto.SHA256
	}else{
		return "",errors.New("不支持该加密方式")
	}

	//获取加密结果
	signBuf := dosign(suit,getByte(data,"UTF-8"),privateKey)


	return "",nil
}
func getByte(data,encoding string) ([]byte,error){
    var dataRune []rune = []rune(data)
	var result []byte
	if encoding =="UTF-8"{
		buf := make([]byte, 2)
        var n int
		for i:=0;i<len(dataRune);{
			n = utf8.EncodeRune(buf,dataRune[i])
			result = append(result,buf...)
			i+=n
		}
		return result,nil

	}else{
		return nil,errors.New("不存在该编码方式")
	}

}