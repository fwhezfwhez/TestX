package test_model

import (
	"fmt"
	"testing"
	"time"
)

func TestUser_ToRepresentation(t *testing.T) {
	u := User{
		Name:      "ft2",
		Age:       9,
		Salary:    1000,
		CreatedAt: time.Now(),
	}

	buf, er := u.ToRepresentation(func(m Block)(func(b Block)Block,[]string){
		//pick :=[]string(nil)
		pick  := []string{"name","age"}
		handler := func(b Block)Block{
			b.Update("question","why?")
			b.Update("location","china")
			b.Update("created_at", u.CreatedAt.Format("2006-01-02 15:04:05"))
			b.Pop("location")
			return b
		}
		return handler,pick
	})
	if er!=nil {
		fmt.Println(er.Error())
		return
	}
	fmt.Println(string(buf))
}
func BenchmarkUser_ToRepresentation(b *testing.B) {
	u := User{
		Name:      "ft2",
		Age:       9,
		Salary:    1000,
		CreatedAt: time.Now(),
	}

	var buf []byte
	var er error
	var pick []string
	var handler func(b Block)Block
	for i := 0; i < b.N; i++ {
		buf, er = u.ToRepresentation(func(m Block)(func(b Block)Block,[]string){
			pick =[]string(nil)
			handler = func(b Block)Block{
				b.Update("question","why?")
				b.Update("location","china")
				b.Update("created_at", u.CreatedAt.Format("2006-01-02 15:04:05"))
				return b
			}
			return handler,pick
		})
		if er!=nil {
			fmt.Println(er.Error())
			return
		}
		fmt.Println(string(buf))
	}
}
