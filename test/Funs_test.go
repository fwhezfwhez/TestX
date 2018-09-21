package test

import "testing"

func TestAdd(t *testing.T) {
	r,er:= (Add(3,3))
	if er !=nil {
		t.Fatal(er.Error())
	}
	t.Log(r)
}
func Add(a,b int)(int,error){
	return a+b,nil
}
func BenchmarkAdd(b *testing.B) {
	for i:=0;i<b.N;i++{
		Add(3,6)
	}

}