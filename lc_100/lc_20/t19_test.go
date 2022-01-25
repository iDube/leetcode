package lc_20

import (
	"fmt"
	"testing"
)

// 19. 删除链表的倒数第 N 个结点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	slow, fast := head, head
	for fast.Next != nil {
		if n > 0 {
			n--
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}

	if n == 1 {
		// 移除的是第一个节点
		return head.Next
	} else {
		slow.Next = slow.Next.Next
		return head
	}
}

func arrayToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	for h, i := head, 1; i < len(nums); i++ {
		h.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		h = h.Next
	}
	return head
}

func Test_removeNthFromEnd(t *testing.T) {
	h := arrayToList([]int{1, 2, 3})
	fmt.Println(removeNthFromEnd(h, 3))
}
