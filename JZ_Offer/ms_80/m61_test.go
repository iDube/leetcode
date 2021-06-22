package ms_80

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// 面试题61. 扑克牌中的顺子

func isStraight(nums []int) bool {
	// 1. 先排序
	sort.Ints(nums)
	fmt.Println(nums)

	zoreCount := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zoreCount++
			continue
		}

		x := nums[i+1] - nums[i]

		if x == 0 {
			return false
		}
		if x == 1 {
			continue
		}
		if x > 1 {
			zoreCount = zoreCount - x + 1
			if zoreCount < 0 {
				return false
			}
		}
	}

	return true
}

func TestIsStraight(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(isStraight([]int{0, 0, 1, 3, 5}), true)
	assertions.Equal(isStraight([]int{0, 0, 1, 2, 5}), true)
	assertions.Equal(isStraight([]int{1, 2, 3, 4, 5}), true)
	assertions.Equal(isStraight([]int{0, 1, 1, 2, 5}), false)
	assertions.Equal(isStraight([]int{0, 0, 3, 1, 2}), true)
}
