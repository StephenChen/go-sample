package base

import "strings"

// Split 把字符串 s 按照给定的分隔符 sep 进行分割，返回字符串切片
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)

	return
}
