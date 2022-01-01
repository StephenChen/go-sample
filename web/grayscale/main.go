package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

var bucketSize = 10

func main() {
	// 判断 hash 算法的均衡度
	var bucketMap = map[uint64]int{}
	for i := 15000000000; i < 15000000000+10000000; i++ {
		hashInt := murmur64_(fmt.Sprint(i)) % uint64(bucketSize)
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
}

func murmur64_(p string) uint64 {
	return murmur3.Sum64([]byte(p))
}
