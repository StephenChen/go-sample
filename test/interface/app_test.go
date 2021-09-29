package _interface

import (
	"github.com/prashantv/gostub"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// 为全局变量 configFile 打桩
	stubs := gostub.Stub(&configFile, "./test.toml")
	defer stubs.Reset() // 测试结束后重置

	data, err := GetConfig()
	if err != nil {
		t.Fatal()
	}

	t.Logf("data: %s\n", data)
}

func TestShowNumber(t *testing.T) {
	stubs := gostub.Stub(&maxNum, 20)
	defer stubs.Reset()

	res := ShowNumber()
	if res != 20 {
		t.Fatal()
	}
}
