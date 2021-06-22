package lc_6000

// 5357. 设计一个支持增量操作的栈

/*

请你设计一个支持下述操作的栈。

实现自定义栈类 CustomStack ：
- CustomStack(int maxSize)：用 maxSize 初始化对象，maxSize 是栈中最多能容纳的元素数量，栈在增长到 maxSize 之后则不支持 push 操作。
- void push(int x)：如果栈还未增长到 maxSize ，就将 x 添加到栈顶。
- int pop()：返回栈顶的值，或栈为空时返回 -1 。
- void inc(int k, int val)：栈底的 k 个元素的值都增加 val 。如果栈中元素总数小于 k ，则栈中的所有元素都增加 val 。

*/

type CustomStack struct {
	size int   // 容量
	top  int   // 栈顶下标
	data []int // 数据
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		size: maxSize,
		top:  -1,
		data: make([]int, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	if this.top+1 < this.size {
		// 有容量
		this.top++
		this.data[this.top] = x
	}
	// 否则，不添加
}

func (this *CustomStack) Pop() int {
	if this.top > -1 {
		// 有元素
		result := this.data[this.top]
		this.top--
		return result
	} else {
		// 题目要求，栈空返回 -1
		return -1
	}
}

func (this *CustomStack) Increment(k int, val int) {
	// 栈底k个元素加val
	for i := 0; i <= this.top && i < k; i++ {
		this.data[i] += val
	}

}

/*
["CustomStack","push","push","pop","push","push","push","increment","increment","pop","pop","pop","pop"]
[[3],          [1],   [2],   [],   [2],   [3],   [4],    [5,100],    [2,100],    [],   [],  [],   []]

["CustomStack","push","push","pop","push","push","push","increment","increment","pop","pop","pop","pop"]
[[3],			[1],  [2],   [],   [2],   [3],   [4],    [5,100],    [2,100],    [],[],[],[]]
输出：
[null,null,null,2,null,null,null,null,null,103,202,101,-1]
预期：
[null,null,null,2,null,null,null,null,null,103,202,201,-1]

*/
