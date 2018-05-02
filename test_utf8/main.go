package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := make([]byte, utf8.UTFMax)

	n := utf8.EncodeRune(b, '好')
	fmt.Printf("%v：%v\n", b, n) // [229 165 189 0]：3


	//n表示所占用的byte字节数
	r, n := utf8.DecodeRune(b)
	fmt.Printf("%c：%v\n", r, n) // 好：3

	s := "大家好"
	for i := 0; i < len(s); {
		r, n = utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%c：%v   ", r, n) // 大：3   家：3   好：3
		i += n
	}
	fmt.Println()

	for i := len(s); i > 0; {
		r, n = utf8.DecodeLastRuneInString(s[:i])
		fmt.Printf("%c：%v   ", r, n) // 好：3   家：3   大：3
		i -= n
	}
	fmt.Println()

	b = []byte("好")
	fmt.Printf("%t, ", utf8.FullRune(b))     // true
	fmt.Printf("%t, ", utf8.FullRune(b[1:])) // true
	fmt.Printf("%t, ", utf8.FullRune(b[2:])) // true
	fmt.Printf("%t, ", utf8.FullRune(b[:2])) // false
	fmt.Printf("%t\n", utf8.FullRune(b[:1])) // false

	b = []byte("大家好")
	fmt.Println(utf8.RuneCount(b)) // 3

	fmt.Printf("%d, ", utf8.RuneLen('A'))          // 1
	fmt.Printf("%d, ", utf8.RuneLen('\u03A6'))     // 2
	fmt.Printf("%d, ", utf8.RuneLen('好'))          // 3
	fmt.Printf("%d, ", utf8.RuneLen('\U0010FFFF')) // 4
	fmt.Printf("%d\n", utf8.RuneLen(0x1FFFFFFF))   // -1

	fmt.Printf("%t, ", utf8.RuneStart("好"[0])) // true
	fmt.Printf("%t, ", utf8.RuneStart("好"[1])) // false
	fmt.Printf("%t\n", utf8.RuneStart("好"[2])) // false

	b = []byte("你好")
	fmt.Printf("%t, ", utf8.Valid(b))     // true
	fmt.Printf("%t, ", utf8.Valid(b[1:])) // false
	fmt.Printf("%t, ", utf8.Valid(b[2:])) // false
	fmt.Printf("%t, ", utf8.Valid(b[:2])) // false
	fmt.Printf("%t, ", utf8.Valid(b[:1])) // false
	fmt.Printf("%t\n", utf8.Valid(b[3:])) // true

	fmt.Printf("%t, ", utf8.ValidRune('好'))        // true
	fmt.Printf("%t, ", utf8.ValidRune(0))          // true
	fmt.Printf("%t, ", utf8.ValidRune(0xD800))     // false  代理区字符
	fmt.Printf("%t\n", utf8.ValidRune(0x10FFFFFF)) // false  超出范围
}
