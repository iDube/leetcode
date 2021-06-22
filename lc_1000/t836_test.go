package lc_1000

// 836. 矩形重叠

// 如果 a,b 连个矩阵重合，a 和 b 完全一样，或者 b 至少有一个角的定点在 a 中

func isRectangleOverlap(rec1 []int, rec2 []int) bool {

	// 题目说明，只考虑矩阵是正放的，这样就不用给点排序了

	// 矩形如果不重叠，从x轴和y轴上看两个矩形就变成了两条线段，这两条线段肯定是不相交的，
	// 也就是说左边的矩形的最右边小于右边矩形的最左边，也就是rec1[2] < rec2[0] || rec2[2] < rec1[0]；
	// y轴同理，下面的矩形的最上边小于上面矩形的最下边，也就是rec1[3] < rec2[1] || rec2[3] < rec1[1]。
	// 因为题目要求重叠算相离，所以加上=，最后取反就行啦~

	return !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0] || rec1[3] <= rec2[1] || rec2[3] <= rec1[1])

}
