package rsa5

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"unicode/utf8"
)

//默认客户端，pkcs8私钥格式，pem编码
func NewDefault(privateKey, publicKey string) (Cipher, error) {
	blockPri, _ := pem.Decode([]byte(privateKey))
	if blockPri == nil {
		return nil, errors.New("private key error")
	}

	blockPub, _ := pem.Decode([]byte(publicKey))
	if blockPub == nil {
		return nil, errors.New("public key error")
	}

	return New(blockPri.Bytes, blockPub.Bytes,PKCS8)
}

func New(privateKey, publicKey []byte, privateKeyType Type) (Cipher, error) {

	priKey, err := genPriKey(privateKey, privateKeyType)
	if err != nil {
		return nil, err
	}
	pubKey, err := genPubKey(publicKey)
	if err != nil {
		return nil, err
	}
	return &pkcsClient{privateKey: priKey, publicKey: pubKey}, nil
}

func genPubKey(publicKey []byte) (*rsa.PublicKey, error) {
	pub, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func genPriKey(privateKey []byte, privateKeyType Type) (*rsa.PrivateKey, error) {
	var priKey *rsa.PrivateKey
	var err error
	switch privateKeyType {
	case PKCS1:
		{
			priKey, err = x509.ParsePKCS1PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
		}
	case PKCS8:
		{
			prkI, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
			priKey = prkI.(*rsa.PrivateKey)
		}
	default:
		{
			return nil, errors.New("unsupport private key type")
		}
	}
	return priKey, nil
}


func getByte(data,encoding string) ([]byte,error){
	var dataRune []rune = []rune(data)
	var result []byte =make([]byte,0)
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