package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

const (
	iv = "abcdefijkloiklkd"
)

func main() {
	fmt.Println(len([]byte(iv)))
	key, e := base64.StdEncoding.DecodeString("1df09d0a1677dd72b8325aec59576e0c")
	if e != nil {
		panic(e)
	}
	es := "OpCoJgs7RrVgaMNDixIvaCIyV2SFDBNLivgkVqtzq2GC10egsn+PKmQ/+5q+chT8xzldLUog2haTItyIkKyvzvmXonBQLIMeq54axAu9c3KG8IhpFD6+ymHocmx07ZKi7eED3t0KyIxJgRNSDkFk5RV1ZP2mSWa7ZgCXXcAbP0RsiUcvhcJfrSwlpsm0E1YJzKpYy429xrEEGvK+gfL+Cw=="
	//es, e := AesEncrypt(str, key)
	//if e != nil {
	//	panic(e)
	//}
	//fmt.Println(es)

	ds, e := AesDecrypt(es, key)
	if e != nil {
		panic(e)
	}
	fmt.Println(ds)
	payloadLen := binary.BigEndian.Uint32(ds[16:20])
	fmt.Println(string(json.RawMessage(ds[20:20+payloadLen])))
}

func AesEncrypt(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS7Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}
func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
