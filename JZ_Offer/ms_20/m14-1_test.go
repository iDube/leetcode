package ms_20

// 剑指 Offer 14- I. 剪绳子

// 数论：
// 1. 任何大于1的数都可由2和3相加组成（根据奇偶证明）
// 2. 因为2*2=1*4，2*3>1*5, 所以将数字拆成2和3，能得到的积最大

func cuttingRope(n int) int {

	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	r := 1

	for n > 4 {
		r *= 3
		n -= 3
	}

	return r * n

}
