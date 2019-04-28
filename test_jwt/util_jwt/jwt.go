package util_jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)
const (
	// SecretKey default
	defaultSecretKey = "zxk$&=*ek$t0u(jfzn^0dfzm7w1vau*r*$%50n@&cq@d7wmoa$"
)
type Jwt struct{
	SecretKey string
}

// new a jwt tool object
func NewJwtTool(secretKey string) Jwt{
	if secretKey == "" {
		return Jwt{
			SecretKey: defaultSecretKey,
		}
	}
	return Jwt{
		SecretKey: secretKey,
	}
}
func (j *Jwt) SetSecretKey(secretKey string) {
	j.SecretKey = secretKey
	return
}

// GenerateJWT generate a jwt string
func (j Jwt)GenerateJWT(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	jwtString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", errors.New(err.Error())
	}
	return jwtString, nil
}

// ValidateJWT validates jwt string
func (j Jwt)ValidateJWT(t string) (*jwt.Token, string) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return token, err.Error()
	}
	return token, ""
}
