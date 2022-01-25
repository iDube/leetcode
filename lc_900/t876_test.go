package lc_900

// 876. 链表的中间结点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	i, j := head, head

	step := 0
	for j.Next != nil {
		if step%2 == 0 {
			i = i.Next
			j = j.Next
		} else {
			j = j.Next
		}
		step++
	}

	return i
}
