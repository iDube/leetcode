package lc_20

import (
	"fmt"
	"strconv"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s string
	for l != nil {
		s = s + strconv.Itoa(l.Val) + " "
		l = l.Next
	}
	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	lResult := &ListNode{
		Val:  0,
		Next: nil,
	}

	point := lResult

	// 进位
	up := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 == nil {
			sum = l2.Val + up
			l2 = l2.Next
		} else if l2 == nil {
			sum = l1.Val + up
			l1 = l1.Next
		} else {
			sum = l1.Val + l2.Val + up
			l1 = l1.Next
			l2 = l2.Next
		}

		if sum >= 10 {
			up = 1
			sum = sum - 10
		} else {
			up = 0
		}
		point.Val = sum

		if l1 != nil || l2 != nil {
			point.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			point = point.Next
		}

	}

	if up != 0 {
		point.Next = &ListNode{
			Val:  up,
			Next: nil,
		}
	}

	return lResult
}

// 测试
func TestAddTwoNumbers(t *testing.T) {

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(l1)
	fmt.Println(l2)

	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3)

}
