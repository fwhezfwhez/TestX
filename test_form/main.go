package main

import "net/http"
type Form struct{
	ViewState string `form`
}
var client = http.Client{}
func main() {
	req,_:= http.NewRequest("POST",)
}
