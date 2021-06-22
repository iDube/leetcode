package ms_40

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题21. 调整数组顺序使奇数位于偶数前面

func exchange(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]%2 == 1 && nums[j]%2 == 0 {
			i++
			j--
		} else if nums[i]%2 == 0 && nums[j]%2 == 0 {
			j--
		} else if nums[i]%2 == 1 && nums[j]%2 == 1 {
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

func Test_exchange(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(exchange([]int{1, 2, 3, 4}), []int{1, 3, 2, 4})
}
