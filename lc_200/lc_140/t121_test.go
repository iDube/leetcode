package lc_140

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 121. 买卖股票的最佳时机

// 暴力穷举, 效率不是最优秀的
func maxProfit(prices []int) int {
	// 边界值处理
	if len(prices) < 2 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}

	mp := 0

	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			mp = max(prices[j]-prices[i], mp)
		}
	}

	// 修正, 貌似没必要
	if mp <= 0 {
		mp = 0
	}
	return mp

}

func Test_maxProfit(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	assertions.Equal(maxProfit([]int{7, 6, 4, 3, 1}), 0)
	assertions.Equal(maxProfit([]int{7, 6, 4, 3, 1, 7, 6, 4, 3, 1}), 6)
	assertions.Equal(maxProfit([]int{1}), 0)
}

// ================================================

// 动态规划
// 记录【今天之前买入的最小值】
// 计算【今天之前最小值买入，今天卖出的获利】，也即【今天卖出的最大获利】
// 比较【每天的最大获利】，取最大值即可
// 转移方程 dp[i] = max(dp[i-1],prices[i] - min(prices[i-1])) 第i天的收益最大值就是 当前价格减去历史最低值
func maxProfit2(prices []int) int {

	// 边界值处理
	if len(prices) < 2 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		} else {
			return b
		}
	}

	mp := 0
	mi := prices[0]

	for i := 0; i < len(prices); i++ {
		mp = max(mp, prices[i]-mi)
		mi = min(mi, prices[i])
	}

	return mp
}

func Test_maxProfit2(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(maxProfit2([]int{7, 1, 5, 3, 6, 4}), 5)
	assertions.Equal(maxProfit2([]int{7, 6, 4, 3, 1}), 0)
	assertions.Equal(maxProfit2([]int{7, 6, 4, 3, 1, 7, 6, 4, 3, 1}), 6)
	assertions.Equal(maxProfit2([]int{1}), 0)
}
