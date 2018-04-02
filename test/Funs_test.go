package test

import "testing"

func TestAdd(t *testing.T) {
	t.Log(Add(3,3))
}

func BenchmarkAdd(b *testing.B) {
	for i:=0;i<b.N;i++{
		Add(3,6)
	}
}