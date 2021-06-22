package ms_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 剑指 Offer 19. 正则表达式匹配

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	// first_match = bool(s and p[0] in {s[0],'.'})
	first_match := len(s) > 0 && (p[0] == s[0] || p[0] == '.')

	// 如果 p 第二个字母是 *
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first_match && isMatch(s[1:], p))
	} else {
		return first_match && isMatch(s[1:], p[1:])
	}

}

func TestIsMatch(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(isMatch("ab", ".*c"), false)
	assertions.Equal(isMatch("mississippi", "mis*is*p*."), false)
	assertions.Equal(isMatch("aa", "a"), false)
	assertions.Equal(isMatch("aa", "a*"), true)
	assertions.Equal(isMatch("aa", ".*"), true)
	assertions.Equal(isMatch("aab", "c*a*b"), true)
	assertions.Equal(isMatch("aaa", "a.a"), true)
	assertions.Equal(isMatch("aaa", "ab*ac*a"), true)
}
