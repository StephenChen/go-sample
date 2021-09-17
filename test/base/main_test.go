package base

import "testing"

// go test -short
func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("short 模式下会跳过该测试用例")
	}
}

// 子测试
func TestXXX(t *testing.T) {
	t.Run("case1", func(t *testing.T) { _ = 1 + 1 })
	t.Run("case2", func(t *testing.T) { _ = 2 + 2 })
	t.Run("case3", func(t *testing.T) { _ = 3 + 3 })
}

// 测试覆盖率
// go test -cover

// 覆盖率相关的记录信息输出文件
// go test -cover -coverprofile=c.out

// 使用浏览器浏览 HTML 报告
// go tool cover -html=c.out
