package ms_60

import (
	"strings"
	"testing"
)

// 面试题58 - I. 翻转单词顺序

func reverseWords(s string) string {
	temp := strings.Split(s, " ")
	result := ""
	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] == "" {
			// 多个连续空格，Split 会切分出一些空字符串
			continue
		}
		result = result + strings.TrimSpace(temp[i]) + " "
	}
	return strings.TrimSpace(result)
}

func TestReverseWords(t *testing.T) {
	reverseWords("a good   example")
}
