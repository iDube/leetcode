package lc_300

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 219. 存在重复元素 II

func containsNearbyDuplicate(nums []int, k int) bool {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j-i <= k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func Test_containsNearbyDuplicate(t *testing.T) {

	assertions := assert.New(t)

	assertions.Equal(containsNearbyDuplicate([]int{99, 99}, 2), true)
	assertions.Equal(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2), false)

}
