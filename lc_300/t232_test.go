package lc_300

// 232. 用栈实现队列

type MyQueue struct {
	StackIn  []int
	StackOut []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		StackIn:  []int{},
		StackOut: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.StackIn = append(this.StackIn, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		// 题目没给出，空时pop出啥，正常情况应该  panic
		return 0
	}
	if len(this.StackOut) == 0 {
		for len(this.StackIn) > 0 {
			this.StackOut = append(this.StackOut, this.StackIn[len(this.StackIn)-1])
			this.StackIn = this.StackIn[:len(this.StackIn)-1]
		}
	}
	// pop
	v := this.StackOut[len(this.StackOut)-1]
	this.StackOut = this.StackOut[:len(this.StackOut)-1]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		// 题目没给出，空时pop出啥，正常情况应该  panic
		return 0
	}
	if len(this.StackOut) == 0 {
		for len(this.StackIn) > 0 {
			this.StackOut = append(this.StackOut, this.StackIn[len(this.StackIn)-1])
			this.StackIn = this.StackIn[:len(this.StackIn)-1]
		}
	}
	// peek
	v := this.StackOut[len(this.StackOut)-1]
	return v
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.StackIn) == 0 && len(this.StackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
