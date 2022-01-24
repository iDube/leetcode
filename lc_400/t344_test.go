package lc_400

import (
	"fmt"
	"testing"
)

// 344. 反转字符串

func reverseString(s []byte) {

	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

}

func Test_reverseString(t *testing.T) {

	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s)
	reverseString(s)
	fmt.Println(s)

}
