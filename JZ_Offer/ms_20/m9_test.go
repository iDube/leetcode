package ms_20

// 面试题09. 用两个栈实现队列

// 先定义一个栈，用链表实现

type StackNode struct {
	Val  int
	Next *StackNode
}

type CStack struct {
	size int
	head *StackNode
}

// 入栈
func (cs *CStack) Push(val int) {
	// 头部追加
	cs.head = &StackNode{
		Val:  val,
		Next: cs.head,
	}
	cs.size++
}

// 出栈
func (cs *CStack) Pop() int {
	if cs.size == 0 {
		// 根据题目要求，没有就返回-1
		return -1
	}
	// 去头
	sn := cs.head
	cs.head = cs.head.Next
	cs.size--
	return sn.Val
}

func (cs *CStack) Size() int {
	return cs.size
}

// 用两个栈实现队列
// 思路：
// 1. 一个栈用于入队，一个栈用于出队;
// 2. 当出队栈为空时，将入队栈内容全部拷贝到出队栈
// 3. 不是并发安全的

type CQueue struct {
	in  *CStack
	out *CStack
}

func Constructor() CQueue {
	return CQueue{
		in: &CStack{
			size: 0,
			head: nil,
		},
		out: &CStack{
			size: 0,
			head: nil,
		},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.out.Size() == 0 && this.in.Size() == 0 {
		// 根据题目要求，没有元素返回-1
		return -1
	}

	// 如果出队栈为空，那么将入队栈拷贝到出队栈内
	if this.out.Size() == 0 {
		for this.in.Size() != 0 {
			this.out.Push(this.in.Pop())
		}
	}
	return this.out.Pop()
}
