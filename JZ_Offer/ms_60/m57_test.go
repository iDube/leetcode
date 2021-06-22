package ms_60

import (
	"fmt"
	"testing"
)

// 面试题57. 和为s的两个数字

func twoSum(nums []int, target int) []int {

	i, j := 0, len(nums)-1

	for i < j {
		if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 13))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))
	fmt.Println(twoSum([]int{10, 26, 30, 31, 47, 60}, 40))
}
