package _14

func splitListToParts(head *ListNode, k int) []*ListNode {
	// 计算链表的长度
	var length, curr = 0, head
	for curr != nil {
		length++
		curr = curr.Next
	}

	width, remainder := length/k, length%k
	res := make([]*ListNode, k)
	curr = head

	for i := 0; i < k; i++ {
		dummy := &ListNode{Val: -1}
		node := dummy
		realWidth := width
		if i < remainder {
			realWidth += 1
		}

		for j := 0; j < realWidth; j++ {
			node.Next = &ListNode{Val: curr.Val}
			node = node.Next
			if curr != nil {
				curr = curr.Next
			}
		}

		res[i] = dummy.Next
	}

	return res
}
