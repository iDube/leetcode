package lc_180

// 栈
type Stack struct {
	size  int
	array []int
}

// 入栈
func (s *Stack) Push(val int) {
	s.array = append(s.array, val)
	s.size++
}

// 出栈
func (s *Stack) Pop() int {
	if s.size == 0 {
		return -1
	}

	sn := s.array[s.size-1]
	s.array = s.array[:s.size-1]
	s.size--
	return sn
}

// Top 返回栈顶数据
func (s *Stack) Top() int {
	if s.size == 0 {
		return -1
	}
	return s.array[s.size-1]
}

func (s *Stack) Size() int {
	return s.size
}

type MinStack struct {
	S Stack
	M Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.S.Push(x)

	if this.M.Size() > 0 {
		if x <= this.M.Top() {
			this.M.Push(x)
		}
	} else {
		this.M.Push(x)
	}
}

func (this *MinStack) Pop() {
	if this.S.Top() == this.M.Top() {
		this.S.Pop()
		this.M.Pop()
	} else {
		this.S.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.S.Top()
}

func (this *MinStack) GetMin() int {
	return this.M.Top()
}
