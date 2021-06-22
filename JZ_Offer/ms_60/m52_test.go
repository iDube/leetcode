package ms_60

import "testing"

// 面试题52. 两个链表的第一个公共节点

// 第一想到的，使用map
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	cache := make(map[*ListNode]bool)

	for headA != nil {
		cache[headA] = true
		headA = headA.Next
	}

	var head *ListNode
	ok := false
	for headB != nil {
		_, ok = cache[headB]
		if ok {
			head = headB
			break
		} else {
			headB = headB.Next
		}
	}
	return head
}

// 另一种，看别人做的，有点辗转相除法的味道
// 设交集链表长c,链表1除交集的长度为a，链表2除交集的长度为b，有
// a + c + b = b + c + a
// 若无交集，则a + b = b + a
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	h1, h2 := headA, headB
	for h1 != h2 {
		if h1 == nil {
			h1 = headB
		} else {
			h1 = h1.Next
		}
		if h2 == nil {
			h2 = headA
		} else {
			h2 = h2.Next
		}
	}
	return h1
}

func Test_getIntersectionNode(t *testing.T) {

	headSame := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: headSame,
		},
	}

	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: headSame,
			},
		},
	}

	getIntersectionNode(headA, headB)
}
