package ms_60

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题50. 第一个只出现一次的字符

func firstUniqChar(s string) byte {

	temp := make(map[byte]int)

	for i := range s {
		temp[s[i]]++
	}

	// map 无序，所以又遍历一遍字符串
	for i := range s {
		if temp[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}

func TestFirstUniqChar(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(firstUniqChar(""), byte(' '))
	assertions.Equal(firstUniqChar("abaccdeffb"), byte('d'))
	assertions.Equal(firstUniqChar("abaccdeff"), byte('b'))
}
