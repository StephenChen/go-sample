package monkey

import (
	"bou.ke/monkey"
	"strings"
	"testing"
)

func TestMyFunc(t *testing.T) {
	// 对 GetInfoByUID 进行打桩
	// 无论传入的 uid 是多少，都返回 &UserInfo{Name: "chenxinyuan"}, nil
	monkey.Patch(GetInfoByUID, func(int64) (*UserInfo, error) {
		return &UserInfo{Name: "chenxinyuan"}, nil
	})
	defer monkey.UnpatchAll()

	ret := MyFunc(123)
	if !strings.Contains(ret, "chenxinyuan") {
		t.Fatal()
	}
}
