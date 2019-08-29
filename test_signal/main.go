package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
		log.Println(<-ch)
		os.Exit(0)
		// do things when catch a close signal
	}()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(1)
	}
}
