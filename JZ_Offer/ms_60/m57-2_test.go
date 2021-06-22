package ms_60

import (
	"fmt"
	"testing"
)

// 面试题57 - II. 和为s的连续正数序列
// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

// 想法：
// 这是遍历嘛，深度优先遍历嘛？
// 连续数列，试试双指针法？

// 双指针法，往后挪动, 貌似有人叫滑动窗口法
func findContinuousSequence(target int) [][]int {
	var result [][]int

	// 检测，小于target -1, 大于 target 1，等于target 0
	check := func(i, j int) int {
		sum := 0
		for ; i <= j; i++ {
			sum += i
		}
		if sum == target {
			return 0
		} else if sum < target {
			return -1
		} else {
			return 1
		}
	}

	buildArray := func(i, j int) []int {
		size := j - i + 1
		r := make([]int, size)
		for t := 0; t < size; t++ {
			r[t] = i
			i++
		}
		return r
	}

	i := 1
	j := i + 1
	// 看了解析，发现终止其实可以改为 target/2+1 ，+1是因为整形 3/2+1 结果才是2
	//for i < target && j < target && i < j {
	end := target/2 + 1
	for i < end && j <= end && i < j {
		temp := check(i, j)
		if temp == 0 {
			// 加起来和结果相等，记录，往后挪
			result = append(result, buildArray(i, j))
			i++
		}
		if temp == -1 {
			// 加起来比结果小，加一个数字，
			j++
		}
		if temp == 1 {
			// 加起来比结果大，减少一个数字
			i++
		}
	}
	return result
}

//
func Test_findContinuousSequence(t *testing.T) {
	fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
	fmt.Println(findContinuousSequence(459))
}
