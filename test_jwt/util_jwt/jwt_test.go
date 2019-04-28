package util_jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	JwtTool.SetSecretKey("HELLO")
	fmt.Println(JwtTool.GenerateJWT(map[string]interface{}{
		"user_id": int(1),
		"version": 1,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	}))
}

func TestValidateJWT(t *testing.T) {
	JwtTool.SetSecretKey("HELLO")
	token, msg := JwtTool.ValidateJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTAyMTM2OTQsInVzZXJfaWQiOjEsInZlcnNpb24iOjF9.ctG6UFunYCcHrNYM1FNSTGk8I04eF1nk4TgogYH1i_8")

	if !token.Valid {
		fmt.Println("valid fail", msg)
		return
	}
	fmt.Println(token.Claims.(jwt.MapClaims)["user_id"])
	fmt.Println(reflect.TypeOf(token.Claims.(jwt.MapClaims)["user_id"]))

	var r map[string]interface{}
	r = token.Claims.(jwt.MapClaims)

	fmt.Println(r)
	fmt.Println(r["user_id"])
	fmt.Println(reflect.TypeOf(r["user_id"]))
}

func TestDecodeJWT(t *testing.T) {
	var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0MTcsInVzZXJfbmFtZSI6IuW8oOWuh-mUiyIsImdyb3VwX2lkIjoyLCJpc19zdXBlcnVzZXIiOmZhbHNlLCJleHAiOjE1NTQ2ODQ2OTN9.mIqLbNSzNGA_ufvZ8zSxoNNEIWffNSRR7zmPw9Ry3xU"
	token, msg :=JwtTool.ValidateJWT(jwtToken)
	if !token.Valid {
		fmt.Println("valid fail", msg)
		return
	}
	var r map[string]interface{}
	r = token.Claims.(jwt.MapClaims)

	fmt.Println(r)
}
