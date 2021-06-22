package lc_40

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 30. 包含min函数的栈

type MinStack struct {
	stack    []int // 栈
	minStack []int // 存放着，将当前push操作后的最小值放在这个栈内
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0, 16),
		minStack: make([]int, 0, 16),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		// 放第一个元素
		this.stack = append(this.stack, x)
		this.minStack = append(this.minStack, x)
	} else {
		if x <= this.minStack[len(this.minStack)-1] {
			this.minStack = append(this.minStack, x)
		}
		this.stack = append(this.stack, x)
	}
}

func (this *MinStack) Pop() {
	if this.Top() == this.Min() {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	// 题目没给出，如果栈内没数据时，给出啥
	// 为了不出错，返回个0
	// 真正工程中，所有的方法，都要加个 err ，来出来各种错误
	if len(this.stack) == 0 {
		return 0
	} else {
		return this.stack[len(this.stack)-1]
	}
}

func (this *MinStack) Min() int {
	if len(this.stack) == 0 {
		return 0
	} else {
		return this.minStack[len(this.minStack)-1]
	}
}

func TestMinStack(t *testing.T) {

	assertions := assert.New(t)

	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	assertions.Equal(obj.Min(), -3)
	obj.Pop()
	assertions.Equal(obj.Top(), 0)
	assertions.Equal(obj.Min(), -2)

}
