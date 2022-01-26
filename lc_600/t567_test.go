package lc_600

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 567. 字符串的排列

func checkInclusion(s1 string, s2 string) bool {

	str1, str2 := []rune(s1), []rune(s2)
	if len(str1) > len(str2) {
		return false
	}

	windowSize := len(str1)
	// 目标词典与动态缓存str2词典
	targetDic, dynamicDic := [26]rune{}, [26]rune{}
	// 初始化词典
	for i := 0; i < windowSize; i++ {
		targetDic[str1[i]-'a']++
		dynamicDic[str2[i]-'a']++
	}

	for i := windowSize; i < len(str2); i++ {
		// 匹配到
		if targetDic == dynamicDic {
			return true
		}
		// 窗口又移动，更新str2动态词典
		dynamicDic[str2[i-windowSize]-'a']--
		dynamicDic[str2[i]-'a']++
	}

	return targetDic == dynamicDic
}

func Test_checkInclusion(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(checkInclusion("ab", "eidbaooo"), true)
	assertions.Equal(checkInclusion("ab", "eidboaoo"), false)
	assertions.Equal(checkInclusion("a", "a"), true)
	assertions.Equal(checkInclusion("a", "b"), false)
	assertions.Equal(checkInclusion("a", "ab"), true)

}
