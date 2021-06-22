package ms_40

// 面试题39. 数组中出现次数超过一半的数字
// 和 t169 一样

// t169 用了 map

// 摩尔投票算法是基于这个事实：
// 每次从序列里选择两个不相同的数字删除掉（或称为“抵消”），
// 最后剩下一个数字或几个相同的数字，就是出现次数大于总数一半的那个。
func majorityElement(nums []int) int {

	// 题目给出，数组长度 >= 1

	result, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
		}
		if result == nums[i] {
			count++
		} else {
			count--
		}
	}

	return result

}
