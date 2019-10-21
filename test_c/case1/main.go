
package main

/*
#include <stdio.h>
*/
import "C"

func main() {
	C.puts(C.CString("Hello, 世界\n"))
}
