package loadbalance

import (
	"fmt"
	"math/rand"
	"time"
)

var endpoints = []string{
	"100.69.62.1:3232",
	"100.69.62.32:3232",
	"100.69.62.42:3232",
	"100.69.62.81:3232",
	"100.69.62.11:3232",
	"100.69.62.113:3232",
	"100.69.62.101:3232",
}

// 洗牌算法，错误：洗牌不均匀。7次均不被选中((6/7)*(6/7))^7=0.34 != 1/7=0.14
func shuffle(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

// fisher-yates
func shuffle1(indexes []int) {
	for i := len(indexes); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
	}
}

func shuffle2(n int) []int {
	b := rand.Perm(n) // 已内置
	return b
}

func resetSeed() {
	rand.Seed(time.Now().UnixNano())
}

func request(params map[string]interface{}) error {
	var indexes = []int{0, 1, 2, 3, 4, 5, 6}
	var err error

	shuffle(indexes)
	maxRetryTimes := 3

	idx := 0
	for i := 0; i < maxRetryTimes; i++ {
		err = apiRequest(params, indexes[idx])
		if err != nil {
			break
		}
		idx++
	}

	if err != nil {
		// logging
		return err
	}

	return nil
}

func apiRequest(params map[string]interface{}, i int) error {
	return nil
}

func init() {
	resetSeed()
}

func LoadBalance() {
	var cnt1 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var sl = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle(sl)
		cnt1[sl[0]]++
	}

	var cnt2 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var sl = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle1(sl)
		cnt2[sl[0]]++
	}

	fmt.Println(cnt1)
	fmt.Println(cnt2)
}
