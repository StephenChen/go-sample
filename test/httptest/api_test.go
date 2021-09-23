package httptest

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestGetResultByAPI(t *testing.T) {
	defer gock.Off() // 测试执行后刷新挂起的 mock

	tests := []struct {
		X   int
		Y   int
		Res int
	}{
		{1, 1, 101},
		{2, 2, 202},
	}

	// mock 请求外部 api 时传参 x=1 返回 100
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 1}).
		Reply(200).
		JSON(map[string]int{"value": 100})

	// mock 请求外部 api 时传参 x=2 返回 200
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 2}).
		Reply(200).
		JSON(map[string]int{"value": 200})

	for _, tt := range tests {
		// 调用业务函数
		res := GetResultByAPI(tt.X, tt.Y)
		// 校验返回结果是否符合预期
		assert.Equal(t, res, tt.Res)
	}

	assert.True(t, gock.IsDone())
}
