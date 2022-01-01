package lock

import (
	"fmt"
	"sync"
)

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	l := Lock{c: make(chan struct{}, 1)}
	l.c <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	select {
	case <-l.c:
		return true
	default:
	}
	return false
}

func (l Lock) UnLock() {
	l.c <- struct{}{}
}

var counter int

func TryLock() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				fmt.Println("lock failed")
				return
			}
			counter++
			fmt.Println("current counter", counter)
			l.UnLock()
		}()
	}
	wg.Wait()
}
