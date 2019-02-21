package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type JR struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

var G *int
func main() {
	var X int
	G = &X
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...
	var cm = make(map[string]JR)
	var tmp string
	for i := 0; i < 100000; i++ {
		tmp = fmt.Sprintf("user_%d", i)
		cm[tmp] = JR{Name: tmp, Data: []byte(`{"kk":9}`)}
	}
	fmt.Println(len(cm))

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
