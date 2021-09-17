package base

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// go test -v
func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

// go test -run=Sep -v
func TestSplitWithComplexSep(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

// 表格驱动测试，例如标准库中 fmt 的测试 TestFlagParser
func TestSplitAll(t *testing.T) {
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected: %#v, got: %#v", tt.want, got)
			}
		})
	}
}

// 并行测试
func TestSplitAllParallel(t *testing.T) {
	// 将 TLog 标记为能够与其他测试并行运行
	t.Parallel()

	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}

	for _, tt := range tests {
		// 重新声明 tt 避免多个 goroutine 中使用了相同的变量
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// 将每个测试用例标记为能够彼此并行运行
			t.Parallel()
			got := Split(tt.input, tt.sep)
			// 使用 testify 简化
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("expected:%#v, got:%#v", tt.want, got)
			//}
			assert.Equal(t, got, tt.want)
		})
	}
}
