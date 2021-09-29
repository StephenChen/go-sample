package monkey

import (
	"bou.ke/monkey"
	"reflect"
	"strings"
	"testing"
)

func TestUser_GetInfo(t *testing.T) {
	var u = &User{
		Name:     "chenxinyuan",
		Birthday: "1994-10-02",
	}

	// 为对象方法打桩
	monkey.PatchInstanceMethod(reflect.TypeOf(u), "CalcAge", func(*User) int {
		return 18
	})

	ret := u.GetInfo()
	if !strings.Contains(ret, "朋友") {
		t.Fatal()
	}
}
