package ms_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题10- II. 青蛙跳台阶问题

// 设跳上 n 级台阶有 f(n) 种跳法，青蛙的最后一步只有两种情况： 跳上 1 级或 2 级台阶。
// 当为 1 级台阶： 剩 n−1 个台阶，此情况共有 f(n−1) 种跳法；
// 当为 2 级台阶： 剩 n−2 个台阶，此情况共有 f(n−2) 种跳法。
// f(n) 为以上两种情况之和，即 f(n)=f(n−1)+f(n−2)
func numWays(n int) int {
	// 边界处理
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1 // 没明白为什么0级太极又一种方式
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		// 题目要求，答案取模
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

func Test_numWays(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(numWays(7), 21)
	assertions.Equal(numWays(2), 2)
}
