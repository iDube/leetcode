package lc_340

import (
	"fmt"
	"testing"
)

// 322. 零钱兑换
// 这题有疑问的：
// 1. coins 是有序的嘛？

// 动态规划问题

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

// 超出时间限制了
func coinChange(coins []int, amount int) int {

	// 边界值处理
	if len(coins) == 0 {
		return -1
	}
	// 出始条件
	if amount == 0 {
		return 0
	}

	// 假设要一个很大的值
	result := 1 << 32

	for _, v := range coins {
		if v == 0 {
			// 处理边界，不处理 coins 里的0
			continue
		}
		if amount >= v {
			result = min(coinChange(coins, amount-v)+1, result)
		}
	}

	if result == 1<<32 || result == 0 {
		// result 没变和 result 等于0都是没找到（因而计算式子里有+1, -1+1=0)
		return -1
	} else {
		return result
	}

}

// 不用递归，把 0 到 amount 全算出来
// 动态规划
// 转移方程：dp[i] = min{dp[i-v] + 1 | v in coins}
func coinChange2(coins []int, amount int) int {

	// 边界值处理
	if len(coins) == 0 {
		return -1
	}

	f := make([]int, amount+1)

	// 出始条件, amount=0 时，就是0
	f[0] = 0

	for i := 1; i <= amount; i++ {
		// 假设要一个很大的值
		f[i] = 1 << 32
		for _, v := range coins {
			if i >= v && f[i-v] != 1<<32 {
				f[i] = min(f[i-v]+1, f[i])
			}

		}
	}

	if f[amount] == 1<<32 {
		f[amount] = -1
	}
	return f[amount]
}

func Test_coinChange(t *testing.T) {
	fmt.Println(coinChange2([]int{1, 2, 5}, 100))
	fmt.Println(coinChange2([]int{2}, 3))
}
