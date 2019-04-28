package test_bigEnd

import "fmt"

func main(){
	var box interface{}
	var a = 1
	var in *int= &a

	box = in
	out := box.(*int)

	*out = 5

	fmt.Println(*in)
}
