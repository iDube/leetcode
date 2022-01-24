package lc_40

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 35. 搜索插入位置

func searchInsert(nums []int, target int) int {

	i, j := 0, len(nums)-1
	for i <= j {

		m := (i + j) / 2
		if target == nums[m] {
			return m
		} else if target > nums[m] {
			i = m + 1
		} else {
			j = m - 1
		}

	}

	return i
}

func Test_searchInsert(t *testing.T) {

	assertions := assert.New(t)
	assertions.Equal(searchInsert([]int{1, 3, 5, 7, 9, 11}, 5), 2)
	assertions.Equal(searchInsert([]int{1, 3, 5, 7, 9, 11}, 6), 3)
	assertions.Equal(searchInsert([]int{1, 3, 5, 7, 9, 11}, 7), 3)
	assertions.Equal(searchInsert([]int{1, 3, 5, 7, 9, 11}, 8), 4)
	assertions.Equal(searchInsert([]int{1}, 2), 1)
	assertions.Equal(searchInsert([]int{1}, 1), 0)
	assertions.Equal(searchInsert([]int{1}, 0), 0)

}
