package server

import (
	"strings"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

/*
   编写一段简单的字符串大写化的微服务
*/
type StringService interface {
	UpperCase(in string) (string, error)
	Count(in string) int
}

type stringService struct {
}

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func (ss *stringService) UpperCase(in string) (string, error) {
	return strings.ToUpper(in),nil
}
func (ss *stringService) Count(in string) int {
	return len(in)
}

func UpperCaseEndPoint(svc StringService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{})(interface{}, error){

	}
}