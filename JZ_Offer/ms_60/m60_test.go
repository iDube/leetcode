package ms_60

import (
	"fmt"
	"math"
	"testing"
)

// 面试题60. n个骰子的点数

// 动态规划问题

// n个骰子，刨出的情况有 6^n 次方个，
// 和 s 为 {n, n+1, n+2, ... 6n-1, 6n} n点出现的次数应该为一次，这时所有骰子都是1点

// dp[n][j] 定义为 投n枚骰子，j点出现的次数：
// dp[1][1] = 1, dp[1][2] = 1, dp[1][3] = 1, dp[1][4] = 1, dp[1][5] = 1, dp[1][6] = 1
// 一枚骰子，1到6点，出现的次数都是1

// 2枚骰子呢？ 点数范围在 2到12，
// dp[2][1] = 0, dp[2][2] = 1, dp[2][3] = 2 ...

//
// dp[n][j] = SUM(dp[n-1][j-i])  i取1到6m, 假设为，先抛出了n-1个，再抛第n个的点数为i
// dp[2][5] = SUM(dp[1][4], dp[1][3], dp[1][2], dp[1][1])
//			= 4

func ts(n int) []float64 {

	// 处理边界值
	if n == 0 {
		return []float64{}
	}

	dp := [][]float64{}

	// 为了逻辑清晰，数组下标从 1 开始，匹配骰子点数
	dp = append(dp, []float64{})
	dp = append(dp, []float64{0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0})

	for i := 2; i <= n; i++ {

		// 第i个骰子，点数有 6i 种情况
		count := 6 * i

		item := make([]float64, count+1)
		for j := i; j <= count; j++ {
			sum := 0.0
			for k := 1; k <= 6 && k <= j; k++ {
				if j-k <= (i-1)*6 {
					sum += dp[i-1][j-k]
				}
			}
			item[j] = sum
		}
		fmt.Println(item)
		dp = append(dp, item)
	}

	result := dp[n]
	fmt.Println(result)
	result = result[n : 6*n]
	s := math.Pow(float64(6), float64(n))
	fmt.Println(s)
	for i := range result {
		result[i] = result[i] / s
	}

	return result
}

func TestTs(t *testing.T) {
	fmt.Println(ts(3))
	fmt.Println(ts(2))
	fmt.Println(ts(1))
}
