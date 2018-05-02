package main

import (
	"testing"

)


var us = User{"FT"}
func BenchmarkAddr(b *testing.B) {
	for i:=0;i<b.N;i++{
		Addr(&us)
	}
	b.Log(us)
}
func BenchmarkReturn(b *testing.B) {
	for i:=0;i<b.N;i++{
		us = Return(us)
	}
	b.Log(us)
}