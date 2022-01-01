package lock

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func ZooKeeper() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}
	fmt.Println("lock success, do business logic")

	time.Sleep(time.Second * 10)

	// do something
	l.Unlock()
	fmt.Println("unlock success, finish business logic")

}
