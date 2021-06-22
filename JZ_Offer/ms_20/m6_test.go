package ms_20

import "testing"

// 面试题06. 从尾到头打印链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 放数组里，数组反转？
// 2. 用栈？
// 3. 用数组反转吧
func reversePrint(head *ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	// 反转数组
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return result
}

func Test_reversePrint(t *testing.T) {

}
