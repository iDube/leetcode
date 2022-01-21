package lc_800

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 704. 二分查找

func search(nums []int, target int) int {

	i, j := 0, len(nums)-1

	for i <= j {
		mid := (j-i)/2 + i
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	// 不存在目标值
	return -1
}

func Test_search(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(search([]int{1}, 1), 0)
	assertions.Equal(search([]int{1}, 0), -1)
	assertions.Equal(search([]int{1, 2}, 1), 0)
	assertions.Equal(search([]int{1, 2}, 2), 1)
	assertions.Equal(search([]int{-1, 0, 3, 5, 9, 12}, 9), 4)
	assertions.Equal(search([]int{-1, 5}, 5), 1)
}
