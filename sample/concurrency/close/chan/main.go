package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, cancel chan bool, i int) {
	defer wg.Done()

	select {
	default:
		fmt.Println("hello", i)
	case <-cancel:
		return
	}
}

func main() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel, i)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
