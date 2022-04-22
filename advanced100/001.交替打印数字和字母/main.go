package main

import (
	"fmt"
	"sync"
)

// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func main() {
	wg := sync.WaitGroup{}

	numCh := make(chan bool)
	chCh := make(chan bool)

	go func() {
		i := 1
		for {
			<-numCh
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			chCh <- true
		}
	}()

	wg.Add(1)

	go func() {
		i := 0
		for {
			<-chCh
			if i == 26 {
				wg.Done()
				return
			}
			fmt.Print(string(rune('a' + i)))
			i++
			fmt.Print(string(rune('a' + i)))
			i++
			numCh <- true
		}
	}()

	numCh <- true
	wg.Wait()
}
