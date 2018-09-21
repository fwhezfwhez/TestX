package main

import (
 "test_X/test_protobuf/Im"
 "github.com/golang/protobuf/proto"
 "log"
)

func main() {
 // 创建一个消息
 test := &Im.Helloworld{
  // 使用辅助函数设置域的值
  Str: "hello!" ,
  //  Id:  321,
  Opt: 1234,
 }

 test.Id = 3244

 // 进行编码
 data, err := proto.Marshal(test)
 if err != nil {
  log.Fatal("marshaling error: ", err)
 }

 // 进行解码
 newTest := &Im.Helloworld{}
 err = proto.Unmarshal(data, newTest)
 if err != nil {
  log.Fatal("unmarshaling error: ", err)
 }

 log.Printf("id:%d;opt:%d;str:%s;",newTest.Id,newTest.Opt,newTest.Str)

 // 测试结果
 if test.String() != newTest.String() {
  log.Fatalf("data mismatch %q != %q", test.String(), newTest.String())
 }
}