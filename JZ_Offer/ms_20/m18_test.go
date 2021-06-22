package ms_20

import "testing"

// 面试题18. 删除链表的节点

/* m6_test.go 中已经定义
type ListNode struct {
	Val  int
	Next *ListNode
}*/

func deleteNode(head *ListNode, val int) *ListNode {
	point := head
	var pointUp *ListNode
	for point != nil {
		if point.Val == val {
			// 删除节点
			if pointUp == nil {
				// 要删除的的是头节点
				head = point.Next
			} else {
				pointUp.Next = point.Next
			}
			// 题目要求，只删除一个
			break
		}
		pointUp = point
		point = point.Next
	}
	return head
}

func ArrayToListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  array[0],
		Next: nil,
	}

	point := head
	for i := 1; i < len(array); i++ {
		point.Next = &ListNode{
			Val:  array[i],
			Next: nil,
		}
		point = point.Next
	}
	return head
}

func Test_ArrayToListNode(t *testing.T) {
	// debug 看吧
	ArrayToListNode([]int{1, 2, 3, 4, 5, 6})
}

func Test_deleteNode(t *testing.T) {
	l := ArrayToListNode([]int{1, 2, 3, 4, 5, 6})
	deleteNode(l, 3)
}
