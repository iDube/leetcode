package lc_300

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 300. 最长上升子序列

// 思路：
// 1. 取第一个数 a[i]，后面比第一个大的数组成一个新的数组 b,
//    那么a的最长子序列就变成了 1 加上 b的最长子序列的后结果的 max。
//    穷举i得到b,继续递归。然后求最大值。是嘛？貌似很复杂
// 2. 找出最大值，最大值将数组分成两部分，
//    假设最大值序号是i f(0...n) = max(f(0 ... i-1)+1,f(i+1 ... n))
//    然后用动态规范的方法做，怎么递归转循环呢？
//	  有两个连续的相等值怎么处理呢？如 [2,3,3,7] 貌似取最大值时，取重复数子的第一个数就好
//	  啥？最大值将数组分成两部分？那么 [2,5,3,4] 呢，错在这里；

// 看别人的
// 动态规划的思路：
//   将 dp 数组定义为：以 nums[i] 结尾的最长上升子序列的长度
//   那么题目要求的，就是这个 dp 数组中的最大者
//   以数组  [10, 9, 2, 5, 3, 7, 101, 18] 为例：
//   dp 的值： 1  1  1  2  2  3  4    4
func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			// 关键步骤，
			if nums[i] > nums[j] {
				// + 1
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 找最大值
	res := 0
	for i := 0; i < size; i++ {
		res = max(res, dp[i])
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLengthOfLIS(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(lengthOfLIS([]int{10, 9, 2, 5, 3, 4}), 3)
	assertions.Equal(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}), 6)
	assertions.Equal(lengthOfLIS([]int{}), 0)
	assertions.Equal(lengthOfLIS([]int{3}), 1)
	assertions.Equal(lengthOfLIS([]int{2, 3}), 2)
	assertions.Equal(lengthOfLIS([]int{3, 2}), 1)
	assertions.Equal(lengthOfLIS([]int{2, 3, 3, 7}), 3)
	assertions.Equal(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), 4)
}
