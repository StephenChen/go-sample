package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for i := range in {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int, 64)

	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	// Ctrl+C exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("quit (%v)\n", <-sig)
}
