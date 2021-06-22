package ms_20

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// 面试题 01.06. 字符串压缩

func compressString(S string) string {

	if S == "" {
		return ""
	}

	count := 0
	last := S[0]
	result := ""
	for i := 0; i < len(S); i++ {
		if last == S[i] {
			count++
		} else {
			result += string(last)
			result += strconv.Itoa(count)
			last = S[i]
			count = 1
		}
	}
	result += string(last) + strconv.Itoa(count)

	if len(result) < len(S) {
		return result
	} else {
		return S

	}

}

func TestCompressString(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(compressString("a"), "a")
	assertions.Equal(compressString("ab"), "ab")
	assertions.Equal(compressString("aa"), "aa")
	assertions.Equal(compressString("aaa"), "a3")
	assertions.Equal(compressString("aabcccccaaa"), "a2b1c5a3")
}
