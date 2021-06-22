package ms_20

import (
	"fmt"
	"math/big"
	"testing"
)

// 面试题10- I. 斐波那契数列
// 求第n项

// 直接for循环吧，会比递归节省空间嘛？
func fib(n int) int {
	// 缓存历史值
	cache := make(map[int]*big.Int)
	cache[0] = big.NewInt(0)
	cache[1] = big.NewInt(1)

	for i := 2; i <= n; i++ {
		cache[i] = big.NewInt(0).Add(cache[i-1], cache[i-2])
	}
	fmt.Println(cache[n])
	r := big.NewInt(0).Mod(cache[n], big.NewInt(1000000007))
	fmt.Println(r)
	return int(r.Int64())
}

func Test_fib(t *testing.T) {
	/*assertions := assert.New(t)
	assertions.Equal(fib(0), 0)
	assertions.Equal(fib(1), 1)
	assertions.Equal(fib(2), 1)
	assertions.Equal(fib(3), 2)
	assertions.Equal(fib(4), 3)
	assertions.Equal(fib(5), 5)
	assertions.Equal(fib(6), 8)
	assertions.Equal(fib(7), 13)
	assertions.Equal(fib(8), 21)*/

	// 94 已经超出uint64了，用big包试试
	fib(92)
	fib(93)
	fib(94)
	fib(95)

}
