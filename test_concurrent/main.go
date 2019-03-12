package main

import "time"

func main() {
	var m = make(map[string]interface{}, 0)
	m["2"] = 2
	go func() {
		for {
			//m ["1"] = 1
			 m["1"] =1
		}
	}()

	go func() {
		for {
			 m["1"] =2
		}
	}()

	time.Sleep(20 * time.Second)
}

