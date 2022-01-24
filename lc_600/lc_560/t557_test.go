package lc_560

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 557. 反转字符串中的单词 III

func reverseWords(s string) string {

	temp := []rune(s)

	l, r, length := 0, 0, len(temp)

	for r < length {
		if temp[r] == ' ' {
			reverse(temp, l, r-1)
			l = r + 1
		} else if r == length-1 {
			reverse(temp, l, r)
		}
		r++
	}

	return string(temp)
}

func reverse(s []rune, l int, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func Test_reverse(t *testing.T) {
	s := []rune{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s)
	reverse(s, 2, 4)
	fmt.Println(s)
}

func Test_reverseWords(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(reverseWords("Let's take LeetCode contest"), "s'teL ekat edoCteeL tsetnoc")
	assertions.Equal(reverseWords(" Let's take LeetCode contest"), " s'teL ekat edoCteeL tsetnoc")
	assertions.Equal(reverseWords("Let's take LeetCode contest "), "s'teL ekat edoCteeL tsetnoc ")

}
