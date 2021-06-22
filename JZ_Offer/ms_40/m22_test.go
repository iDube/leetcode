package ms_40

// 面试题22. 链表中倒数第k个节点

// 用例额外的空间
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 边界值处理
	if head == nil || k < 1 {
		return head
	}

	var cache []*ListNode
	for head != nil {
		cache = append(cache, head)
		head = head.Next

	}

	if len(cache) >= k {
		return cache[len(cache)-k]
	} else {
		return cache[0]
	}
}

// 快慢指针
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	p1 := head
	p2 := head
	for p1 != nil && k >= 0 {
		if k > 0 {
			p1 = p1.Next
			k--
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return p2
}
