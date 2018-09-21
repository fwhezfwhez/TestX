package rsa

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"testing"
	"runtime"
)

//var  privateKey2 = []byte(`
//-----BEGIN PRIVATE KEY-----
//MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCfU8v4BUr81SKm
///H0ahbdQZjEpO8nMyk+xuYSatHwnU4//m47R+4G2YB4Z6PHsJi4+ScfJpQutFhKr
//FwTXZ6TDqLvaqZDDkJq5G271g+PmrzFp7f40/E9m0qjeL64RJra0rZql23dvPW4v
//VomMRgRcoPOn0YWVp+M6T5PaFgE4M8dh4lMZz57gVwOdd08F99Z92f3QgZtEjI+/
//EXvMenXxb/aRofNkt+Wdk2ELJ6MIP0d9UU5v3WgLuuNv5QnQYzj/RMr8GD+wrDYi
//NQJxsaTmE/OEJggsumhD4eYY5YlRy2EIN504cujYVKU1wOSZgq9oJCynGR0aPuQW
//x58IHxEtAgMBAAECggEAHfEFd8qm2PTE2lTAvec7F+TcgD84IUAz0dZnURtx6YIO
//oZ5+LH/zVG6juYLJU/Oo5RPAc+iMVS68u2JMCp7zm8Ft7B3JkrbuHLNHGuR6Q7PQ
//uXN8PkDcOxqDmZ2kPJzl4PZvBZRE0abdug+tMatGzpGAuJzrWcB/N0oVIvrXp9Pn
//Oqfo/Y5nxmpOFCImJppIS3AL1pftNtQZo9G15CPHDYtpUbXPtD2MjjW4OLxKuPRo
//HSwUgo6LW9XSwNXfcuK+lbzLL0BhlWD9IV/+yCEUEblN87yxxfhpQFaAhXj5W+B3
//YsMOZuK93+XMOpYmw8EpUDMObOnvwb0NSHUrV2RUAQKBgQDTojlnNS1e7+tjPzFt
//OhGPj1uCBPAEIeHAcnPgd80bEiujxMLCnGaAvmnTrMu4Xo0e5fAP4F7R6UD+IUsf
//r3CAAu7CadQ49TW+SovAvciy9AZuSVVIwynu6QdYgFyPKe1LZYAEq5k+mB1Vh5q0
//RoxMNAA5pGYKg8+4MmmsJi7X7QKBgQDAunCOqIiH128bs/1VRIhDpzuRW5Qr/SRb
//O2saVg5RSHnO/nGT2OuxSTTkc8yrx7qd9SmAxXl5kR238DhMOQOnRBomldmVtAJu
//JgrdQyt0wXfeQVQqshqCUaE/xhEbpSCdbPSZbKZZdplV0y6O5vXIhxw+1qAvXLcx
//w46s3R92QQKBgQClQ+ejywkVPDILHMwSSehwvThufkCYWYUbbcVDowpOe5AMoZid
//tNju7MNjg2rLHTsCx/kBzOr+7THNwl4R7kTiEmg09cO+fu5rHXepGgtig+GJukaZ
//PZ6/bMZJvGOLgOhHmomwG/jdwpgVtIGBCh6BW5JZcSImT+ykIOoYfvDRuQKBgCgw
//OHxnBGFfORoLxE3dhpSk8LT05cbueIBVuZW6UC3+8PeK82AjIbLMUy04QHupoG6D
//yu3BP/1rl0jd3L94PBzLBLD7Gm4vJTqW0DknYo5sMXS1JrnofcKjBv7nbHXZTx3E
//tJSxpVaOdpcA/HpsCuCP3AH2e1yk9sZ3wu6lBYSBAoGACYM60j1CVRNSZxUNRgiw
//fWzS69qI1eezPc7xQEganpVBI9SZcTNp1kpDKmQikXJ4Yb5XWn12HCY/sFeBW6Su
//3ruNqxvg1XiUPbH6A6nxd5B3QX0mS9+wDm6ONysPLRdKbfFO0mdP4CeyuGPdvDIM
//XP4dJdLhMUL4pcJLI0B7gBE=
//-----END PRIVATE KEY-----
//`)

var privateKey2 = []byte(`
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

var publicKey = []byte(`
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
var cipher Cipher

func init() {
	client, err := NewDefault(string(Private), string(Public))

	if err != nil {
		fmt.Println(Location(),err)
	}

	cipher = client
}

func Test_DefaultClient(t *testing.T) {

	cp, err := cipher.Encrypt([]byte("测试加密解密"))
	if err != nil {
		t.Error(Location(),err)
	}
	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(Location(),err)
	}
	pp, err := cipher.Decrypt(ppBy)

	fmt.Println(string(pp))
}

func Test_Sign_DefaultClient(t *testing.T) {

	src := `attach=附加信息&body=测试购买商品&charset=UTF-8&device_info=AND_SDK
&mch_app_id=com.wwl.tmgp.sgame&mch_app_name=玩玩乐&mch_create_ip=127.0.0.1&mch_id=175510359638&
nonce_str=1522309983261&notify_url=http://xxx/weixinwap-pay/testPayResult&out_trade_no=1522309969277
&service=pay.weixin.wappay&sign_type=RSA_1_256&total_fee=1&version=2.0`

	signBytesRaw, err := cipher.Sign([]byte(src), crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
	returnSignStr := base64.StdEncoding.EncodeToString(signBytesRaw)
	fmt.Println(returnSignStr )

	signBytesRawDe,err := base64.StdEncoding.DecodeString(returnSignStr)
	if err!=nil{
		t.Fatal(err)
	}

	//sign := hex.EncodeToString(signBytesRawDe)
	//
	//
	//signB, err := hex.DecodeString(sign)

	errV := cipher.Verify([]byte(src),signBytesRawDe, crypto.SHA256,"")
	if errV != nil {
		t.Fatal(errV)
	}
	fmt.Println("verify success")
}


func Location() string {
	_, f, line, _ := runtime.Caller(1)
	return fmt.Sprintf("[%s%d]", f, line)
}
