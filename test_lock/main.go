package main

import (
	"runtime"
	"sync"
	"time"
)

type SafeLock struct {
	M        *sync.RWMutex
	hasRLock bool
	hasLock  bool
}

func NewSafeLock() *SafeLock {
	return &SafeLock{
		M:        &sync.RWMutex{},
		hasRLock: false,
		hasLock:  false,
	}
}

func (sl *SafeLock) Lock() {
	if !sl.hasLock {
		sl.hasLock = true
		sl.M.Lock()
	}
}

func (sl *SafeLock) Unlock() {
	if sl.hasLock {
		sl.M.Unlock()
		sl.hasLock = false
	}
}

func (sl *SafeLock) RLock() {
	if !sl.hasRLock {
		sl.hasRLock = true
		sl.M.RLock()
	}
}

func (sl *SafeLock) RUnlock() {
	if sl.hasRLock {
		sl.M.RUnlock()
		sl.hasRLock = false
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sl := NewSafeLock()
	for i := 0; i < 100; i++ {
		go func() {
			sl.Lock()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			sl.Unlock()
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			sl.RLock()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			sl.RUnlock()
		}()
	}
	time.Sleep(20 * time.Second)
}
