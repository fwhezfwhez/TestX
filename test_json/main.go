package main

import (
	"fmt"
)

type C struct {
	In interface{} `json:"in"`
}

type D struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func f(arg interface{}) {
	switch v := arg.(type) {
	case float32, float64:
		fmt.Println(v == 0.)
	case int, int32, int64:
		fmt.Println(v == 0)
	}
}

type User struct {
	Age int
}

func X() (user *User) {
	user.Age = 1
	return
}

func main() {
	var a = make([]string, 0, 5)
	fmt.Println(len(a))
	//var i float32 = 12.3
	//fmt.Println(float32(i)<float32(5.3))
	////var j time.Time
	//var k struct{int}
	//vT := reflect.TypeOf(k)
	//fmt.Println(vT.Field(0).Type.String())
	//
	//fmt.Println(len(strings.Split("",":")))
	//L:
	//for i:=0;i<9;i++{
	//	fmt.Print(i,"")
	//	for j:=0;j<9;j++{
	//		if j==1 {
	//			fmt.Print(j,"")
	//			continue L
	//		}
	//		fmt.Print(j,"")
	//	}
	//	if i==1{
	//		break
	//	}
	//}
	//fmt.Println(fmt.Sprintf("%v",[]interface{}{
	//	1,"2",time.Now(),
	//}))
	//
	//var i float32
	//f(i)
	//
	//m:=make(map[string]string,5)
	//sr:=make([]int,0,5)
	//m["1"]="_1"
	//m["2"]="_2"
	//m["3"]="_3"
	//m["4"]="_4"
	//m["5"]="_5"
	//fmt.Println(len(m))
	//m["6"]="_6"
	//fmt.Println(len(m))
	//
	//sr=append(sr, 1)
	//sr=append(sr, 2)
	//fmt.Println(len(sr),cap(sr))
	//sr=append(sr, 3)
	//sr=append(sr, 4)
	//sr=append(sr, 5)
	//fmt.Println(len(sr),cap(sr))
	//sr=append(sr, 6)
	//sr=append(sr, 7)
	//fmt.Println(len(sr),cap(sr))

}

//func main() {
//	//c := C{}
//	//c.In = D{"ft",5}
//	//b, er := json.Marshal(c)
//	//if er != nil {
//	//	fmt.Println(er.Error())
//	//	return
//	//}
//	//fmt.Println(string(b))
//	//
//	//c.In = [2]D{
//	//	{"ft1",7}, {"ft2",8},
//	//}
//	//b2, er := json.Marshal(c)
//	//if er != nil {
//	//	fmt.Println(er.Error())
//	//	return
//	//}
//	//fmt.Println(string(b2))
//	//
//	//
//	//rs := Query(10,[3]D{
//	//	{"ft1",7}, {"ft2",8},{"ft3",9},
//	//})
//	//b3, er := json.Marshal(rs)
//	//if er != nil {
//	//	fmt.Println(er.Error())
//	//	return
//	//}
//	//fmt.Println(string(b3))
//	var f float32
//	fmt.Println(f==0)
//	in := 0.
//	var tmp interface{} = float32(in)
//	fmt.Println("float 0==0:", in == 0)
//	fmt.Println("float -> interface{} -> float", tmp.(float32) == 0)
//	switch v := tmp.(type) {
//	case float32:
//		fmt.Println("float -> interface -.type-> float", v == 0)
//	}
//}

func Query(count int, in interface{}) interface{} {
	tmp := struct {
		Count   int         `json:"count"`
		Results interface{} `json:"results"`
	}{
		Count:   count,
		Results: in,
	}
	return tmp
}
