package test
func Add(a ...int) int{
	sum := 0
	for _,v:=range a {
		sum = sum+v
	}
	return sum
}