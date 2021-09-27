package database

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestDoSomethingWithRedis(t *testing.T) {
	// mock 一个 redis server
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	// 准备数据
	s.Set("chen", "xinyuan")
	s.SAdd(KeyValidWebsite, "chen")

	// 连接 mock 的 redis server
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	// 调用函数
	ok := DoSomethingWithRedis(rdb, "chen")
	if !ok {
		t.Fatal()
	}

	// 可手动检查 redis 中的值是否符合预期
	if got, err := s.Get("blog"); err != nil || got != "https://xinyuan" {
		t.Fatalf("'blog' has the wrong value")
	}

	// 也可使用帮助工具检查
	s.CheckGet(t, "blog", "https://xinyuan")

	// 过期检查
	s.FastForward(5 * time.Second)
	if s.Exists("blog") {
		t.Fatalf("'blog' should not have existed anymore")
	}
}
