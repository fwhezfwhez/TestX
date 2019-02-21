package httpResponse

import (
	"fmt"
	"golang.org/x/protobuf/proto"
	"test_X/go-micro-user-validate/service/example_user/proto"
	"testing"
)

func TestMalshal(t *testing.T) {
	buf, e := proto.Marshal(&user.Response{Token:"xxx", HttpResponse: &user.HttpResponse{StatusCode:400, Body:"hello"}})
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}

	var rs = &Wrap{}
	e = proto.Unmarshal(buf, rs)
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(rs.HttpResponse.Body, rs.HttpResponse.StatusCode)
}
