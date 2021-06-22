package ms_60

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题59 - II. 队列的最大值

// 思路：
// 1. 首先肯定要满足队列的限制，先进先出； -> 且还要有个查看最大值的功能；
// 2. 关键是对数据结构的选择；-> 队列用头尾双指针链表方便做，用两个链表试试? -> 一个链表保存数据，一个链表做当前（及历史）最大值
// 3. 或者用切片做？也是两个切片，一个保存数据，一个用来做当前最大值
// 4. 链表加一个堆？

type MaxQueue struct {
	dataLength, maxValLength int
	data, maxVal             []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		dataLength:   0,
		maxValLength: 0,
		data:         nil,
		maxVal:       nil,
	}
}

func (this *MaxQueue) Max_value() int {
	if this.maxValLength == 0 {
		if this.dataLength > 0 {
			return this.data[0]
		} else {
			return -1
		}
	} else {
		return this.maxVal[0]
	}
}

func (this *MaxQueue) Push_back(value int) {
	// 从切片末尾入值
	this.data = append(this.data, value)
	this.dataLength++

	// 放入保存当前最大值的切片
	if this.maxValLength == 0 {
		this.maxVal = append(this.maxVal, value)
		this.maxValLength++
	} else {
		for i := this.maxValLength - 1; i >= 0; i-- {
			if i == 0 && value > this.maxVal[i] {
				this.maxVal[0] = value
				break
			} else if value > this.maxVal[i] {
				this.maxVal = this.maxVal[0:i]
				this.maxValLength--
			} else {
				this.maxVal = append(this.maxVal, value)
				this.maxValLength++
				break
			}
		}
	}

}

func (this *MaxQueue) Pop_front() int {
	value := -1
	if this.dataLength > 0 {
		// 从切片头出一个
		value = this.data[0]
		this.data = this.data[1:]
		this.dataLength--
		// 看出的这个是不是当前最大值，是的话，保存当前最大值的切片也从头出一个
		if this.Max_value() == value {
			this.maxVal = this.maxVal[1:]
			this.maxValLength--
		}

	}
	return value
}

func Test_MaxQueue(t *testing.T) {
	mq := Constructor()
	assertions := assert.New(t)
	assertions.Equal(mq.Max_value(), -1)
	//mq.Push_back(2)
	mq.Push_back(1)
	mq.Push_back(2)
	assertions.Equal(mq.Max_value(), 2)
	assertions.Equal(mq.Pop_front(), 1)
	assertions.Equal(mq.Max_value(), 2)
	mq.Push_back(3)
	mq.Push_back(1)
	assertions.Equal(mq.Max_value(), 3)
	assertions.Equal(mq.Pop_front(), 2)
	assertions.Equal(mq.Pop_front(), 3)
	assertions.Equal(mq.Max_value(), 1)
}
