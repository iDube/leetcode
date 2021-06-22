package ms_60

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题53 - II. 0～n-1中缺失的数字

// 变形的二分，如果 right - left + 1 不等于子串的长度，说明缺失的在这个子串里
func missingNumber(nums []int) int {

	i, j := 0, len(nums)-1

	for i <= j {
		m := (i + j) / 2
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return i
}

func TestMissingNumber(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(missingNumber([]int{0}), 1)
	assertions.Equal(missingNumber([]int{1}), 0)
	assertions.Equal(missingNumber([]int{0, 1, 3}), 2)
	assertions.Equal(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}), 8)

}
