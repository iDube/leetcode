package lc_1100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1071. 字符串的最大公因子

// 求最大公约数, 辗转相除，递归法
func gcb2(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcb(b, a%b)
	}
}

// 求最大公约数, 辗转相除，循环法
func gcb(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func TestGcb(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gcb(6, 3), 3)
	assertions.Equal(gcb(3, 6), 3)
	assertions.Equal(gcb(1, 0), 0)

	assertions.Equal(gcb2(6, 3), 3)
	assertions.Equal(gcb2(3, 6), 3)
	assertions.Equal(gcb2(1, 0), 0)
}

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}
	if str1+str2 == str2+str1 {
		return str1[0:gcb(len(str1), len(str2))]
	} else {
		return ""
	}
}

func Test_gcdOfStrings(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(gcdOfStrings("ABCABC", "ABC"), "ABC")
	assertions.Equal(gcdOfStrings("ABABAB", "ABAB"), "AB")
	assertions.Equal(gcdOfStrings("LEET", "CODE"), "")
	assertions.Equal(gcdOfStrings("", ""), "")
	assertions.Equal(gcdOfStrings("A", ""), "")
	assertions.Equal(gcdOfStrings("AAA", "AAA"), "AAA")

}
