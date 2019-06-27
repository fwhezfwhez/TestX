package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFlags(log.Llongfile| log.LstdFlags)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World, %v\n", time.Now())
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Println(s.ListenAndServe())
	}()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	log.Println(<-ch)


	log.Println("server shutdown")
	// Stop the service gracefully.
	log.Println(s.Shutdown(nil))

	// Wait gorotine print shutdown message
	time.Sleep(time.Second * 5)
	log.Println("graceful done.")
}
