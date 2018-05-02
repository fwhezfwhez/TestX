package main

import (
	"crypto/sha256"
	"crypto"
	"fmt"
	"os"
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"crypto/x509"
	"errors"
	"encoding/hex"
)

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
var privateKey = []byte(formatKey("MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCfU8v4BUr81SKm/H0ahbdQZjEpO8nMyk+xuYSatHwnU4//m47R+4G2YB4Z6PHsJi4+ScfJpQutFhKrFwTXZ6TDqLvaqZDDkJq5G271g+PmrzFp7f40/E9m0qjeL64RJra0rZql23dvPW4vVomMRgRcoPOn0YWVp+M6T5PaFgE4M8dh4lMZz57gVwOdd08F99Z92f3QgZtEjI+/EXvMenXxb/aRofNkt+Wdk2ELJ6MIP0d9UU5v3WgLuuNv5QnQYzj/RMr8GD+wrDYiNQJxsaTmE/OEJggsumhD4eYY5YlRy2EIN504cujYVKU1wOSZgq9oJCynGR0aPuQWx58IHxEtAgMBAAECggEAHfEFd8qm2PTE2lTAvec7F+TcgD84IUAz0dZnURtx6YIOoZ5+LH/zVG6juYLJU/Oo5RPAc+iMVS68u2JMCp7zm8Ft7B3JkrbuHLNHGuR6Q7PQuXN8PkDcOxqDmZ2kPJzl4PZvBZRE0abdug+tMatGzpGAuJzrWcB/N0oVIvrXp9PnOqfo/Y5nxmpOFCImJppIS3AL1pftNtQZo9G15CPHDYtpUbXPtD2MjjW4OLxKuPRoHSwUgo6LW9XSwNXfcuK+lbzLL0BhlWD9IV/+yCEUEblN87yxxfhpQFaAhXj5W+B3YsMOZuK93+XMOpYmw8EpUDMObOnvwb0NSHUrV2RUAQKBgQDTojlnNS1e7+tjPzFtOhGPj1uCBPAEIeHAcnPgd80bEiujxMLCnGaAvmnTrMu4Xo0e5fAP4F7R6UD+IUsfr3CAAu7CadQ49TW+SovAvciy9AZuSVVIwynu6QdYgFyPKe1LZYAEq5k+mB1Vh5q0RoxMNAA5pGYKg8+4MmmsJi7X7QKBgQDAunCOqIiH128bs/1VRIhDpzuRW5Qr/SRbO2saVg5RSHnO/nGT2OuxSTTkc8yrx7qd9SmAxXl5kR238DhMOQOnRBomldmVtAJuJgrdQyt0wXfeQVQqshqCUaE/xhEbpSCdbPSZbKZZdplV0y6O5vXIhxw+1qAvXLcxw46s3R92QQKBgQClQ+ejywkVPDILHMwSSehwvThufkCYWYUbbcVDowpOe5AMoZidtNju7MNjg2rLHTsCx/kBzOr+7THNwl4R7kTiEmg09cO+fu5rHXepGgtig+GJukaZPZ6/bMZJvGOLgOhHmomwG/jdwpgVtIGBCh6BW5JZcSImT+ykIOoYfvDRuQKBgCgwOHxnBGFfORoLxE3dhpSk8LT05cbueIBVuZW6UC3+8PeK82AjIbLMUy04QHupoG6Dyu3BP/1rl0jd3L94PBzLBLD7Gm4vJTqW0DknYo5sMXS1JrnofcKjBv7nbHXZTx3EtJSxpVaOdpcA/HpsCuCP3AH2e1yk9sZ3wu6lBYSBAoGACYM60j1CVRNSZxUNRgiwfWzS69qI1eezPc7xQEganpVBI9SZcTNp1kpDKmQikXJ4Yb5XWn12HCY/sFeBW6Su3ruNqxvg1XiUPbH6A6nxd5B3QX0mS9+wDm6ONysPLRdKbfFO0mdP4CeyuGPdvDIMXP4dJdLhMUL4pcJLI0B7gBE=",1,false))
//var  privateKey = []byte(`
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

func main() {
	fmt.Println("privateKey len:",len(privateKey))

	fmt.Println("privateKey2 len:",len(privateKey2))
	fmt.Println(string(privateKey))
	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader
	//message := []byte("hello rsa")
	//message := []byte("body=绿地303观光&charset=UTF-8&device_info=iOS_SDK&mch_app_id=https://www.qinglong365.com/&mch_app_name=青龙移动广告&mch_create_ip=111.75.193.25&mch_id=102525560637&nonce_str=ACnrUlcDryPiltyxa9m6DFtyTopmGrqf&notify_url=https://market.qinglong365.comv1/POST/VXWAP/notifyUrl&out_trade_no=D201804241524565524713447&service=pay.weixin.wappay&sign_type=RSA_1_256&total_fee=1&version=2.0")

//	message := []byte(`attach=附加信息&body=测试购买商品&charset=UTF-8&device_info=AND_SDK
//&mch_app_id=com.wwl.tmgp.sgame&mch_app_name=玩玩乐&mch_create_ip=127.0.0.1&mch_id=175510359638&
//nonce_str=1522309983261&notify_url=http://xxx/weixinwap-pay/testPayResult&out_trade_no=1522309969277
//&service=pay.weixin.wappay&sign_type=RSA_1_256&total_fee=1&version=2.0`)
	message := []byte(`attach=附加信息&body=测试购买商品&charset=UTF-8&device_info=AND_SDK
&mch_app_id=com.wwl.tmgp.sgame&mch_app_name=玩玩乐&mch_create_ip=127.0.0.1&mch_id=175510359638&
nonce_str=1522309983261&notify_url=http://xxx/weixinwap-pay/testPayResult&out_trade_no=1522309969277
&service=pay.weixin.wappay&sign_type=RSA_1_256&total_fee=1&version=2.0`)
	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).
	hashed := sha256.Sum256(message)

	//block,_:= pem.Decode(privateKey) //将密钥解析成私钥实例
	//if er!=nil {
	//	fmt.Println("xx   yyy   zzz",er)
	//	return
	//}
	//if block == nil {
	//	fmt.Println(errors.New("private key error!"))
	//	return
	//}
	//fmt.Println("private block.Byte", block.Bytes)
	priv, err := x509.ParsePKCS8PrivateKey(privateKey) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, err := rsa.SignPKCS1v15(rng, priv.(*rsa.PrivateKey), crypto.SHA256, hashed[:])
	//signature, err := rsa.SignPKCS1v15(rng, priv, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return
	}

	fmt.Printf("Signature: %x\n", signature)

	Decode(message,signature)
//	fmt.Println("Signature len", len(signature))
//	fmt.Println("src len", len(`kB7cUpgVlDovsW8ZxwhofBIBzAItDqepibJBWnvbkEa23woPm0VdUlLpR4LtAPxuma8745iN8vM9n/OOyEtoD50kHhM7PgyoKb
//	AgHOHHAB/Da01D53I2dRmUg93B8cdJbBC7u5TQGG4D+xW7cJ9/MfaeOu9KvU/3//jmTQ1HIz3EuPL6+41vgA+WMiFSN7W8ku4Eh
//	S64Ok3JJ7VetYXX9oSMvOoSSBAdOKHmQJC21i6XilTZAwCiwV+sEJtM9RNf/+2zqndaOXHfBy29R4uSDaAndPhMnndbVYn6RIvn
//	gMH6jTHNIhH2/5LlxDcAJmq8AnuPE2K9dXSBmPu39XKshw==`))
//	fmt.Println("src private key len", len(`MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCfU8v4BUr81SKm
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
//XP4dJdLhMUL4pcJLI0B7gBE=`))
//	fmt.Println("my private key len", len(`MIIEpQIBAAKCAQEApwOQnZ34IZHKQrl8/sgMvn5TlHNNR/+tAK8EU6OEH7n7NCRs
//dgzjSxua8GFtAsoXdzDUM8iCXC+/3cxbSQa3PW8r7sgq5hWZUhQuw0zJFXtXxjLE
//tmNu1o92pS/C5AGpqNR4Qwa3gNeNUxDYQSIIhy9pmGRFV49ZkrN0/3JRAMlECfPv
//R2wC6XMR7hOEfA3k5uudDscXUcAaISACVQAbkbhvx6H+1NKvcb6+m/CdzLNAo6jJ
//YGAePCtN1W9SCYL0FtNc2HfBZ2mK4aik7t6h4UE/yENJPuRhGQEXWDNylJZBueqn
//PIvoQ3zWIKLLUsbYAFoOf/11ikkak5nFRHVDiwIDAQABAoIBAGWglrRCdsW+mAwI
//INZMVuzno+Y2TzVbkCNVXFWkr3Y6znAJJfKjnv+KGbriHdpPrP27ObUU8rYz3BWU
//D4wdSQ1aA6q2JNyDEhvO6jGvHME0n9Gb/PVbBgwLmk+kA3yIwntYZqGqiakeAQIl
//wuHHWtcf86pmgYdpVEEfcDcskbvewRi//yc3OHnO53xUo5EoYLHY98dsSyP3rbvH
//pg0pWLntdyW0+62TmyquivS9L1rgGyYK4sE4cITm41LSdKZI1lPUIU0xjGtyQqUe
//oeN+rtlsb+mBc6Qd5awHR4lslxIOzYnrw+i+cmnwyj0rrSQT8bueYvLnFFLI8MrW
//xl3n/5kCgYEA0j+iUivrFjAgja7iCkpif3DAV5Toq/v7j++l3FV6lOuz/pF36ych
//N4ZGsPl+h+g7j2GfkTG3l3lVanAdkZ1YKbEpta6MrbV8bfMCySABLQXbg+0n5F9k
//KhpcIrhBSS4QU+HuAb88q1sSvlndaPqOzfDBrG/zOKZ0PIaKL8AuOR0CgYEAy1ty
//y9ovAivHM1J4UPahZvmibLPeKKySgJsroI8TlP6AWxUuOKng1scXi8mDaTJldgtV
///zwjSWBb56+RWIahnDE52FqKN4Cow88Z4HErg69o4jMOn/dxGtiVfNCNj1FDSquW
//l3QQMr5SLEfDbJq63kbg5kteOmMjxWvOz5Ap9scCgYEAjtZEhObpc82GkDs7vhoD
//COmlAttbp1wt7/+0AxjfisUHJ+/UNKIE4yAKbrlRyStxK1v7eDz6qaH48bLxo7ft
//6YU+/Pt0/57IxjnOrq1bnybXl61K6NAV0LTP25aqY/kzhol7lRxDL3dUKJYM4gE3
//e5oROgFkkBwl4+jTgNVjkdECgYEAnI/TWc6o0msR+hlMOEkpCja3LilsqLP0Nr9A
//X2TIFrW57GXqKAXh7HlrT1vdqGf8ZV1k4BRUtnFRUJaCuD0uvynG9yL9tq1/QQF5
//UqrObZvmi//zCoVetuR4cpe95NETlbF1RzU8I1UTvrOhaUPXCrpW+/aanSkmqR/P
//IU5EgJECgYEAp3t8YYnARC+kpBVl68szpsgpU8BKGXBEIFUIuV87hvmNDlcqc7nf
//P0DfbJF/qokC0LQ8Zuj6KnN0ooM/z9iAYDWZpn+SQHbQvq9uXTWAK2LuI0t45uP3
//L99yOKf4+gNKlKVyDml/Y40B1lu9TXDdCtLEp7v/7Hjq05Rljw98v0o=`))
//	fmt.Println("src last line len",len("XP4dJdLhMUL4pcJLI0B7gBE="))
//	fmt.Println("my last line len", len("L99yOKf4+gNKlKVyDml/Y40B1lu9TXDdCtLEp7v/7Hjq05Rljw98v0o="))


}

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
		var publicHeader = "\n-----BEGIN PRIVATE KEY-----\n"
		var publicTail = "-----END PRIVATE KEY-----\n"
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
			//fmt.Println(len(*temp)-1)
			key = key[i+1:]
			split(key,temp)
			break
		}
	}
}

func Decode(msg []byte,signature []byte){
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		//return nil, errors.New("public key error")
		fmt.Println(errors.New("public key error"))
	}
	fmt.Println("public block.Byte", block.Bytes)
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		fmt.Println("pub err",err)
	}
	//pub := pubInterface.(*rsa.PublicKey)

	signatureDe, _ := hex.DecodeString(string(signature))
	hashed := sha256.Sum256(msg)

	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signatureDe)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return
	}
	fmt.Println("验证通过")

}