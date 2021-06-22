package lc_6000

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 5368. 找出数组中的幸运数

func findLucky(arr []int) int {

	cache := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		cache[arr[i]]++
	}

	m := -1
	for k, v := range cache {
		if k == v {
			m = max(m, k)
		}
	}

	return m

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestFindLucky(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(findLucky([]int{2, 2, 3, 4}), 2)
	assertions.Equal(findLucky([]int{1, 2, 2, 3, 3, 3}), 3)
	assertions.Equal(findLucky([]int{2, 2, 2, 3, 3}), -1)
	assertions.Equal(findLucky([]int{5}), -1)
	assertions.Equal(findLucky([]int{7, 7, 7, 7, 7, 7, 7}), 7)
}
