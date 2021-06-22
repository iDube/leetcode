package lc_1200

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1160. 拼写单词

// 要查找 chars ，用 map 做吧

func countCharacters(words []string, chars string) int {

	result := 0

	for i := 0; i < len(words); i++ {

		// 用map的话，每次都要初始化，看到别人用数组，数组可以方便的实现 copy, 直接创建个26字母的数组
		chche := make(map[int32]int)
		// 将 chars 放入 map
		for _, v := range chars {
			chche[v]++
		}

		isSpell := true
		for _, v := range words[i] {
			if chche[v] > 0 {
				chche[v]--
			} else {
				isSpell = false
			}
		}
		if isSpell {
			result += len(words[i])
		}
	}

	return result

}

func TestCountCharacters(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"), 6)
	assertions.Equal(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"), 10)
	assertions.Equal(countCharacters([]string{}, "welldonehoneyr"), 0)
	assertions.Equal(countCharacters([]string{"hello", "world", "leetcode"}, "hello"), 5)
	assertions.Equal(countCharacters([]string{"hello", "world", "leetcode"}, ""), 0)
}
