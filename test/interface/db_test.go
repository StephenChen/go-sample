package _interface

import (
	"github.com/golang/mock/gomock"
	"test/interface/mocks"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	// 创建 gomock 控制器，记录后续操作信息
	ctrl := gomock.NewController(t)
	// Finish checks to see if all the methods that were expected to be called
	// were called. It should be invoked for each Controller. It is not idempotent
	// and therefore can only be invoked once.
	//
	// New in go1.14+, if you are passing a *testing.T into NewController function you no
	// longer need to call ctrl.Finish() in your test methods.
	defer ctrl.Finish()

	// 调用 mockgen 生成代码中的 NewMockDB 方法
	m := mocks.NewMockDB(ctrl)
	// 打桩（stub）
	// 当传入 Get 函数的参数为 chenxinyuan 时返回 1 和 nil
	m.EXPECT().Get(gomock.Eq("chenxinyuan")).Return(1, nil).Times(1)
	//m.EXPECT().Get(gomock.Not(gomock.Eq("tr"))).Return(10, nil)
	//m.EXPECT().Get(gomock.Any()).Return(20, nil)
	//m.EXPECT().Get(gomock.Any()).Do(func(key string) {
	//	t.Logf("input key is %v\n", key)
	//})
	//m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
	//	t.Logf("input key is %v\n", key)
	//	return 10, nil
	//})
	//m.EXPECT().Get(gomock.Nil()).Return(-1, nil)

	// 调用GetFromDB函数时传入上面的mock对象m
	if v := GetFromDB(m, "chenxinyuan"); v != 1 {
		t.Fatal()
	}

	// 指定顺序
	gomock.InOrder(
		m.EXPECT().Get("1"),
		m.EXPECT().Get("2"),
		m.EXPECT().Get("3"),
	)
	// 按顺序调用
	GetFromDB(m, "1")
	GetFromDB(m, "2")
	GetFromDB(m, "3")
}
