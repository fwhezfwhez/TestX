package main

import (
  "fmt"
  "net/http"
  "regexp"
)

func main(){
   rsp,er:= http.Get("http://localhost:8090/scheme/")
   if er!=nil {
     fmt.Println(er.Error())
     return
   }
   fmt.Printf(rsp.Proto)

  re, _ := regexp.Compile(pat)
  re.FindAll()
}
