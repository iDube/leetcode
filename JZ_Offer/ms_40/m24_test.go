package ms_40

// 面试题24. 反转链表

func reverseList(head *ListNode) *ListNode {
	p1 := head
	if p1 == nil {
		// 没有节点
		return nil
	}
	p2 := p1.Next
	if p2 == nil {
		// 只有一个节点
		return p1
	}
	p3 := p2.Next

	for true {
		if *p1 == *head {
			p1.Next = nil
		}
		p2.Next = p1

		if p3 != nil {
			p1 = p2
			p2 = p3
			p3 = p3.Next
		} else {
			break
		}
	}
	return p2
}
