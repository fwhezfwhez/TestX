package main


import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)


// 加密
//func RsaEncrypt(origData []byte) ([]byte, error) {
//	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
//	if block == nil {
//		return nil, errors.New("public key error")
//	}
//	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
//	if err != nil {
//		return nil, err
//	}
//	pub := pubInterface.(*rsa.PublicKey)
//	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData) //RSA算法加密
//}


//// 解密
//func RsaDecrypt(ciphertext []byte) ([]byte, error) {
//	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
//	if block == nil {
//		return nil, errors.New("private key error!")
//	}
//	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
//	if err != nil {
//		return nil, err
//	}
//	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext) //RSA算法解密
//}

func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext) //RSA算法解密
}

func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData) //RSA算法加密
}
////私钥
//var privateKey = []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
//7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
//Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
//AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
//ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
//XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
///jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
//IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
//4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
//DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
//9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
//DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
//AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
//-----END RSA PRIVATE KEY-----
//`)

//var privateKey = []byte(formatKey("MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQABAoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaMZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI89KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrwDPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWOAQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O",1,false))
//var publicKey = []byte(formatKey("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB",1,true))
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
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEApwOQnZ34IZHKQrl8/sgMvn5TlHNNR/+tAK8EU6OEH7n7NCRs
dgzjSxua8GFtAsoXdzDUM8iCXC+/3cxbSQa3PW8r7sgq5hWZUhQuw0zJFXtXxjLE
tmNu1o92pS/C5AGpqNR4Qwa3gNeNUxDYQSIIhy9pmGRFV49ZkrN0/3JRAMlECfPv
R2wC6XMR7hOEfA3k5uudDscXUcAaISACVQAbkbhvx6H+1NKvcb6+m/CdzLNAo6jJ
YGAePCtN1W9SCYL0FtNc2HfBZ2mK4aik7t6h4UE/yENJPuRhGQEXWDNylJZBueqn
PIvoQ3zWIKLLUsbYAFoOf/11ikkak5nFRHVDiwIDAQABAoIBAGWglrRCdsW+mAwI
INZMVuzno+Y2TzVbkCNVXFWkr3Y6znAJJfKjnv+KGbriHdpPrP27ObUU8rYz3BWU
D4wdSQ1aA6q2JNyDEhvO6jGvHME0n9Gb/PVbBgwLmk+kA3yIwntYZqGqiakeAQIl
wuHHWtcf86pmgYdpVEEfcDcskbvewRi//yc3OHnO53xUo5EoYLHY98dsSyP3rbvH
pg0pWLntdyW0+62TmyquivS9L1rgGyYK4sE4cITm41LSdKZI1lPUIU0xjGtyQqUe
oeN+rtlsb+mBc6Qd5awHR4lslxIOzYnrw+i+cmnwyj0rrSQT8bueYvLnFFLI8MrW
xl3n/5kCgYEA0j+iUivrFjAgja7iCkpif3DAV5Toq/v7j++l3FV6lOuz/pF36ych
N4ZGsPl+h+g7j2GfkTG3l3lVanAdkZ1YKbEpta6MrbV8bfMCySABLQXbg+0n5F9k
KhpcIrhBSS4QU+HuAb88q1sSvlndaPqOzfDBrG/zOKZ0PIaKL8AuOR0CgYEAy1ty
y9ovAivHM1J4UPahZvmibLPeKKySgJsroI8TlP6AWxUuOKng1scXi8mDaTJldgtV
/zwjSWBb56+RWIahnDE52FqKN4Cow88Z4HErg69o4jMOn/dxGtiVfNCNj1FDSquW
l3QQMr5SLEfDbJq63kbg5kteOmMjxWvOz5Ap9scCgYEAjtZEhObpc82GkDs7vhoD
COmlAttbp1wt7/+0AxjfisUHJ+/UNKIE4yAKbrlRyStxK1v7eDz6qaH48bLxo7ft
6YU+/Pt0/57IxjnOrq1bnybXl61K6NAV0LTP25aqY/kzhol7lRxDL3dUKJYM4gE3
e5oROgFkkBwl4+jTgNVjkdECgYEAnI/TWc6o0msR+hlMOEkpCja3LilsqLP0Nr9A
X2TIFrW57GXqKAXh7HlrT1vdqGf8ZV1k4BRUtnFRUJaCuD0uvynG9yL9tq1/QQF5
UqrObZvmi//zCoVetuR4cpe95NETlbF1RzU8I1UTvrOhaUPXCrpW+/aanSkmqR/P
IU5EgJECgYEAp3t8YYnARC+kpBVl68szpsgpU8BKGXBEIFUIuV87hvmNDlcqc7nf
P0DfbJF/qokC0LQ8Zuj6KnN0ooM/z9iAYDWZpn+SQHbQvq9uXTWAK2LuI0t45uP3
L99yOKf4+gNKlKVyDml/Y40B1lu9TXDdCtLEp7v/7Hjq05Rljw98v0o=
-----END RSA PRIVATE KEY-----
`)
func main() {
	fmt.Println(string(publicKey))
	data, err := RsaEncrypt([]byte(`attach=附加信息&body=测试购买商品&charset=UTF-8&device_info=AND_SDK&mch_app_id=com.wwl.tmgp.sgame&mch_app_name=
玩玩乐&mch_create_ip=127.0.0.1&mch_id=175510359638&nonce_str=1522309983261&notify_url=http://xxx/weixinwap-pay/testPay
Result&out_trade_no=1522309969277&service=pay.weixin.wappay&sign_type=RSA_1_256&total_fee=1&version=2.0`),publicKey) //RSA加密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA加密", string(data))
	origData, err := RsaDecrypt(data,privateKey) //RSA解密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA解密", string(origData))
}

//key是否具有头尾换行不交由程序判断
//keyType=1,为key增加头尾，并每间隔64位换行
//ifPublic true 为公钥， false为私钥
//keyType=0,不变
func formatKey(key string, keyType int, ifPublic bool) string {
	if keyType == 0 {
		return key
	}
	if ifPublic {
		var publicHeader = "\n-----BEGIN PUBLIC KEY-----\n"
		var publicTail = "-----END PUBLIC KEY-----\n"
		var temp string
		split(key,&temp)
		return publicHeader+temp+publicTail
	}else{
		var publicHeader = "\n-----BEGIN RSA PRIVATE KEY-----\n"
		var publicTail = "-----END RSA PRIVATE KEY-----\n"
		var temp string
		split(key,&temp)
		return publicHeader+temp+publicTail
	}
}

func split(key string,temp *string){
	if len(key)<=64 {
		*temp = *temp+key+"\n"
	}
	for i:=0;i<len(key);i++{
		if (i+1)%64==0{
			*temp = *temp+key[:i+1]+"\n"
			fmt.Println(len(*temp)-1)
			key = key[i+1:]
			split(key,temp)
			break
		}
	}
}