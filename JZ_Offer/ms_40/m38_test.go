package ms_40

import (
	"fmt"
	"testing"
)

// 剑指 Offer 38. 字符串的排列

func permutation(s string) []string {

	res := p_dfs([]string{}, []byte{}, s)

	// res 去重
	temp := make(map[string]int)
	for i, _ := range res {
		temp[res[i]] = 0
	}
	keys := make([]string, 0, len(temp))
	for k := range temp {
		keys = append(keys, k)
	}
	return keys
}

func p_dfs(res []string, r []byte, s string) []string {

	if len(s) == 0 {
		res = append(res, string(r))
	} else {
		for i, _ := range s {
			res = p_dfs(res, append(r, s[i]), s[:i]+s[i+1:])
		}
	}
	return res
}

func Test_permutation(t *testing.T) {
	fmt.Println(permutation("abc"))
	fmt.Println(permutation("abb"))
}
