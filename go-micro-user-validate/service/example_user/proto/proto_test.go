package user

import (
	"fmt"
	"golang.org/x/protobuf/proto"
	"testing"
)

func TestMalshal(t *testing.T) {
	buf,e := proto.Marshal(&Response{Token:"tokenxxxxx.xxxx.xxx",HttpResponse:&HttpResponse{StatusCode: 400, Body:"hello"}})
	if e!=nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}

	var rs = &Wrap{}
	e = proto.Unmarshal(buf, rs)
	if e!=nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(rs.HttpResponse.Body, rs.HttpResponse.StatusCode)
}
