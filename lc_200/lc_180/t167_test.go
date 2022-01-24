package lc_180

// 167. 两数之和 II - 输入有序数组

// 双指针
func twoSum(numbers []int, target int) []int {

	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			// 结果，不是从0开始算的
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}
