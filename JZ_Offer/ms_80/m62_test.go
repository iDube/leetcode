package ms_80

// 面试题62. 圆圈中最后剩下的数字

// https://blog.csdn.net/u011500062/article/details/72855826
// 约瑟夫环
// f(N,M)=(f(N−1,M)+M)%N
// p = (p+m)%i

func lastRemaining(n int, m int) int {

	p := 0
	for i := 2; i <= n; i++ {
		p = (p + m) % i
	}
	return p

}
