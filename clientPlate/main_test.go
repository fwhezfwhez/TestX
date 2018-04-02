package main

import "testing"

func BenchmarkMediumList(b *testing.B) {
	for i:=0;i<b.N;i++{
		MediumList()
	}
}