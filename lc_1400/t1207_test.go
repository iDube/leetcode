package lc_1400

import (
	"fmt"
	"testing"
)

/**
1207. 独一无二的出现次数

给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
*/

func uniqueOccurrences(arr []int) bool {
	// 如果数组为空或只有一个元素，那么每个数字出现的次数一定是独一无二的
	if len(arr) < 2 {
		return true
	}

	// 统计各数字的出现个数
	temp := make(map[int]int)
	for _, v := range arr {
		temp[v] = temp[v] + 1
	}

	// 分析是否存在重复
	temp2 := make(map[int]bool)
	for _, v := range temp {
		if _, ok := temp2[v]; ok {
			return false
		} else {
			temp2[v] = true
		}
	}

	return true
}

func Test_uniqueOccurrences(t *testing.T) {

	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr))

}
