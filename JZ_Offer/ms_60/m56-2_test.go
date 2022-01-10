package ms_60

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 剑指 Offer 56 - II. 数组中数字出现的次数 II

func singleNumber(nums []int) int {

	res := nums[0]

	for i := range nums {
		if res == nums[i] {
			res = nums[i+1]
		}
	}

	return 0

}

func Test_singleNumber(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(singleNumber([]int{3, 4, 3, 3}), 4)
}
