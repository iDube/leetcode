package lc_20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 3. 无重复字符的最长子串

// 超时了
func lengthOfLongestSubstring2(s string) int {

	str := []rune(s)

	// size 滑动窗口大小，窗口越来越大
	for size := len(str); size > 0; size-- {
		l := 0
		r := l + size - 1
		for r < len(str) {
			if !hasRepeat(str[l : r+1]) {
				return size
			} else {
				r++
				l++
			}
		}
	}

	return 0

}

func hasRepeat(str []rune) bool {
	temp := make(map[rune]bool)
	for i := 0; i < len(str); i++ {
		if temp[str[i]] == true {
			return true
		} else {
			temp[str[i]] = true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {

	intMax := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	str := []rune(s)

	temp := make(map[rune]int)
	max, start := 0, 0

	for end := 0; end < len(str); end++ {
		ch := str[end]
		if _, ok := temp[ch]; ok {
			start = intMax(temp[ch]+1, start)
		}
		max = intMax(max, end-start+1)
		temp[ch] = end
	}
	return max

}

func Test_lengthOfLongestSubstring(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(lengthOfLongestSubstring("a"), 1)
	assertions.Equal(lengthOfLongestSubstring("aa"), 1)
	assertions.Equal(lengthOfLongestSubstring("abab"), 2)
	assertions.Equal(lengthOfLongestSubstring("ababc"), 3)
	assertions.Equal(lengthOfLongestSubstring("abcabc"), 3)
	assertions.Equal(lengthOfLongestSubstring("abcabcbb"), 3)
	assertions.Equal(lengthOfLongestSubstring("bbbbb"), 1)
	assertions.Equal(lengthOfLongestSubstring("pwwkew"), 3)
	assertions.Equal(lengthOfLongestSubstring(""), 0)

}
