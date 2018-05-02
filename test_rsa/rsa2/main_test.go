package xrsa

import (
	"testing"
	"bytes"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var public = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApwOQnZ34IZHKQrl8/sgM
vn5TlHNNR/+tAK8EU6OEH7n7NCRsdgzjSxua8GFtAsoXdzDUM8iCXC+/3cxbSQa3
PW8r7sgq5hWZUhQuw0zJFXtXxjLEtmNu1o92pS/C5AGpqNR4Qwa3gNeNUxDYQSII
hy9pmGRFV49ZkrN0/3JRAMlECfPvR2wC6XMR7hOEfA3k5uudDscXUcAaISACVQAb
kbhvx6H+1NKvcb6+m/CdzLNAo6jJYGAePCtN1W9SCYL0FtNc2HfBZ2mK4aik7t6h
4UE/yENJPuRhGQEXWDNylJZBueqnPIvoQ3zWIKLLUsbYAFoOf/11ikkak5nFRHVD
iwIDAQAB
-----END PUBLIC KEY-----`
var private = `
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
`

var publicKey *bytes.Buffer = bytes.NewBufferString(public)
var privateKey *bytes.Buffer = bytes.NewBufferString(private)
var xrsa *XRsa

func TestCreateKeys(t *testing.T) {
	err := CreateKeys(publicKey, privateKey, 2048)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestNewXRsa(t *testing.T) {
	var err error
	err = CreateKeys(publicKey, privateKey, 2048)
	xrsa, err = NewXRsa(publicKey.Bytes(), privateKey.Bytes())
	if err != nil {
		t.Error(err.Error())
	}
}

func TestEncryptDecrypt(t *testing.T) {
	xrsa,err := NewXRsa(publicKey.Bytes(), privateKey.Bytes())
	if err != nil {
		t.Fatal(err.Error())
	}
	data := "Estimates of the number of languages中国 in the world vary between 5,000 and 7,000. However, any precise estimate depends on a partly arbitrary distinction between languages and dialects. Natural languages are spoken or signed, but any language can be encoded into secondary media using auditory, visual, or tactile stimuli – for example, in whistling, signed, or braille. This is because human language is modality-independent. Depending on philosophical perspectives regarding the definition of language and meaning, when used as a general concept, language may refer to the cognitive ability to learn and use systems of complex communication, or to describe the set of rules that makes up these systems, or the set of utterances that can be produced from those rules. All languages rely on the process of semiosis to relate signs to particular meanings. Oral, manual and tactile languages contain a phonological system that governs how symbols are used to form sequences known as words or morphemes, and a syntactic system that governs how words and morphemes are combined to form phrases and utterances."
	encrypted, err := xrsa.PublicEncrypt(data)
	if err != nil {
		t.Fatal(err.Error())
	}

	decrypted, err := xrsa.PrivateDecrypt(encrypted)
	if err != nil {
		t.Fatal(err.Error())
	}

	if string(decrypted) != data {
		t.Fatal(fmt.Sprintf("Faild assert \"%s\" equals \"%s\"", decrypted, data))
	}
}

func TestSignVerify(t *testing.T) {
	data := "Hello, World"
	sign, err := xrsa.Sign(data)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = xrsa.Verify(data, sign)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCrossLanguage(t *testing.T) {
	var data = make(map[string] string)
	pubKey, err := ioutil.ReadFile("../../test/pub.pem")
	if err != nil {
		t.Fatal(err.Error())
	}
	priKey, err := ioutil.ReadFile("../../test/pri.pem")
	if err != nil {
		t.Fatal(err.Error())
	}
	testData, err := ioutil.ReadFile("../../test/data.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = json.Unmarshal(testData, &data)
	if err != nil {
		t.Fatal(err.Error())
	}

	rsa2, err := NewXRsa(pubKey, priKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	decrypted, err := rsa2.PrivateDecrypt(data["encrypted"])
	if err != nil {
		t.Fatal(err.Error())
	}
	if string(decrypted) != data["data"] {
		t.Fatal(fmt.Sprintf("Faild assert \"%s\" equals \"%s\"", decrypted, data))
	}

	err = rsa2.Verify(data["data"], data["sign"])
	if err != nil {
		t.Fatal(err.Error())
	}
}

