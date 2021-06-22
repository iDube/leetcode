package ms_40

import (
	"testing"
)

// 面试题25. 合并两个排序的链表

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	p1 := l1
	p2 := l2

	for p2 != nil {
		if p1.Val <= p2.Val {
			if p1.Next == nil || p2.Val < p1.Next.Val {
				p1Next := p1.Next
				p2Next := p2.Next
				p1.Next = p2
				p2.Next = p1Next
				p1 = p1.Next
				p2 = p2Next
			} else {
				p1 = p1.Next
			}
		} else {
			p2Next := p2.Next
			p2.Next = p1
			p1 = p2
			l1 = p1
			p2 = p2Next
		}
	}

	return l1

}

func TestMergeTwoLists(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	b := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	mergeTwoLists(a, b)
}
