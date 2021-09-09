package once

import "sync"

// sync.Once
// 看看 sync.Once 的实现
type singleton struct{}

var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
