package ms_60

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题42. 连续子数组的最大和

// 思路
// 要求时间复杂度为O(n)。那么就只能循环一次了
// 求最值问题，动态规划嘛？
// 数组上，滑动窗口嘛？
// 转换方程: dp[i] = max{dp[i-1], nums[i], ldp[i-1]+nums[i] }, ldp[i] = max{ldp[i-1]+nums[i], nums[i]}
// dp[i] 当前子序列最大和, nums[i] 当前新加入的一个值，ldp[i-1] 包含数组内最后一个值的子序列和最大值
// 化简下  ldp[i] = max { ldp[i-1]+nums[i], nums[i]}， dp[i] = max{dp[i-1], ldp[i] }
func maxSubArray(nums []int) int {

	// 边界值问题
	length := len(nums)
	if len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}

	dp := make([]int, length)
	ldp := make([]int, length)

	dp[0] = nums[0]
	ldp[0] = nums[0]

	for i := 1; i < length; i++ {
		ldp[i] = max(ldp[i-1]+nums[i], nums[i])
		dp[i] = max(dp[i-1], ldp[i])
	}

	return dp[length-1]
}

func Test_maxSubArray(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)

	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2}))
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2}))
}

// 别人的另一个解法，改变了原先的数组，把 nums 直接当 dp 数组了，
func maxSubArray2(nums []int) int {
	n := len(nums)
	maxSum := nums[0]

	max := func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}

	// 如果当前值小于0，
	// 重新开始(全局最大值更新)
	for i := 1; i < n; i++ {
		// 更新当前的最大值
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		// 更新全局的最大值
		maxSum = max(nums[i], maxSum)
	}

	return maxSum

}
