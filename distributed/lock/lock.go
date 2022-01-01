package lock

import (
	"fmt"
	"sync"
)

var lockCounter int
var l sync.Mutex

func InProcess() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			lockCounter++
			l.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(lockCounter)
}
