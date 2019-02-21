package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "./test_prof/myprof.prof", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	for i:=1;i<20;i++{
		log.Println(Add(i,1))
		time.Sleep(1 *time.Second)
	}
	fmt.Println(ToOne(1000))
}

func Add(a,b int) int{
	return a+b
}

func ToOne(n int) int{
	if n == 1 {
		return n
	}
	return ToOne(n-1)
}
