package ms_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*面试题 17.16. 按摩师
一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。
在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。
给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。

注意：本题相对原题稍作改动

示例 1：
输入： [1,2,3,1]
输出： 4
解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。

示例 2：
输入： [2,7,9,3,1]
输出： 12
解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。

示例 3：
输入： [2,1,4,5,3,1,1,3]
输出： 12
解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。*/

// 动态规划问题吧？求最值
// f(n) 长为 n 的输入序列求得的总分钟数
// f(n+1) = f(n) + A[n] 或 f(n)
// 什么情况下加A[n]呢，在没加 A[n-1]的情况下，或者A[n]远岛屿a[n]
// 什么时候不加A[n]呢，在已经加了A[n-1]的情况下

// 递推方程：dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])

func massage(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestMassage(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(massage([]int{1}), 1)
	assertions.Equal(massage([]int{1, 2}), 2)
	assertions.Equal(massage([]int{1, 2, 3}), 4)
	assertions.Equal(massage([]int{1, 2, 3, 1}), 4)
}
