package test_umeng

import (
	"fmt"
	"testing"
)

func TestYoumengOpenApiSignature(t *testing.T) {
	sign :=YoumengOpenApiSignature("param2/1/com.umeng.uapp/umeng.uapp.getLaunches/6297101", map[string]string{
		"startDate":"2018-10-29",
		"endDate":"2018-10-30",
		"appkey":"5b3ecaf7f43e4845c2000168",
	},"ulGpLNDJsw")
	fmt.Println(sign)
}

func TestYoumengOpenApiSignature2(t *testing.T) {
	sign :=YoumengOpenApiSignature("param2/1/system/currentTime/1000000", map[string]string{
		"b":"2",
		"a":"1",
	},"test123")
	fmt.Println(sign)
}

