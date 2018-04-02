package main

import (
	"testing"
)

func TestToken_JwtGenerator(t *testing.T) {
	token := GetToken()
	token.AddHeader("typ", "JWT").AddHeader("alg", "HS256")

	token.AddPayLoad("userName", "admin").AddPayLoad("role", "admin")
	jwt, _, err := token.JwtGenerator("hello")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("签名是:",jwt)
}

func TestToken_Decode(t *testing.T) {
	token := GetToken()
	p, h, hs, err := token.Decode("eyJyb2xlIjoiYWRtaW4iLCJ1c2VyTmFtZSI6ImFkbWluIn0=.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.2s7mpsHJAsp9JLcZWGJ/91KhthCUsDvIvFE0eScg+cM=")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("解出的payload:",p)
	t.Log("解出的header",h)
	t.Log("截出的HS256段",hs)
}

func TestToken_IsLegal(t *testing.T) {
	token := GetToken()
	legal, err := token.IsLegal("eyJyb2xlIjoiYWRtaW4iLCJ1c2VyTmFtZSI6ImFkbWluIn0=.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.2s7mpsHJAsp9JLcZWGJ/91KhthCUsDvIvFE0eScg+cM=", "hello")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("是否合法：",legal)

}
