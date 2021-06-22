package lc_180

import "testing"

// 169. 多数元素

// 排序，确定中间那个，就是多余 n/2 的
// 能不能不完全排序完，就确定了中间的那个？貌似不能

// 使用map
func majorityElement2(nums []int) int {
	tag := len(nums) / 2
	cache := make(map[int]int)
	for _, v := range nums {
		cache[v]++
	}
	for k, v := range cache {
		if v > tag {
			return k
		}
	}
	return -1
}

// 摩尔投票法
func majorityElement(nums []int) int {
	length := len(nums)
	if length == 0 {
		// 题目要求不足已看出，空数组如何处理
		return -1
	}
	if length == 1 || length == 2 {
		return nums[0]
	}

	for i := 0; i < length; i++ {
		count := 0
		v := nums[i]
		for j := 0; j < length; j++ {
			if v == nums[j] {
				count++
			} else {
				count--
			}
		}
		if count > 0 {
			// 题目规定，一定又数量大于 n/2 的
			return nums[i]
		}
	}
	return 0
}

func Test_majorityElement(t *testing.T) {

}
