package lc_500

// 409. 最长回文串

func longestPalindrome(s string) int {

	// 放数组里，然后去构建，只返回长度的话，不需要真的构建，计数就行了

	char := make([]int, 58)
	for _, v := range s {
		char[v-'A']++
	}

	ans := 0
	for _, v := range char {
		// 记偶数
		ans += v - (v & 1)
	}

	// 如果最终的长度小于原字符串的长度，说明里面某个字符出现了奇数次，那么那个字符可以放在回文串的中间，所以额外再加一。
	if ans < len(s) {
		return ans + 1
	} else {
		return ans
	}

}
