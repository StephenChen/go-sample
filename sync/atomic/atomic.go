package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// atomic.AddUint64
var total uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		atomic.AddUint64(&total, uint64(i))
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)

	wg.Wait()
	fmt.Println(total)
}

// atomic.StoreUint32 and atomic.LoadUint32
type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

// atomic.Value
var config atomic.Value    // 保存配置信息
var requests []interface{} // 请求

func loadConfig() interface{} {
	return nil
}

func do() {
	// 初始化配置信息
	config.Store(loadConfig())

	// 启动后台线程，加载更新配置信息
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()

	// 处理请求的工作线程使用最新的配置信息
	for i := 0; i < 10; i++ {
		go func() {
			for _ = range requests {
				_ = config.Load()
				// do...
			}
		}()
	}
}
