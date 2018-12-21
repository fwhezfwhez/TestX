package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	var wg = sync.WaitGroup{}
	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go Map(i, i, &wg)
	}
	wg.Wait()
	fmt.Println("耗时:",time.Now().Sub(t1).String())

}

func TestSyncMap(t *testing.T) {
	var wg = sync.WaitGroup{}
	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go SyncMap(i, i, &wg)
	}
	wg.Wait()
	fmt.Println("耗时:",time.Now().Sub(t1).String())
}
