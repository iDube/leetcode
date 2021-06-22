package ms_60

// 面试题58 - II. 左旋转字符串

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[0:n]
}
