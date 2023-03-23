package _14

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &ListNode{Next: head}
	prev, first, second := dummyNode, head, head.Next

	for second != nil {
		next := second.Next
		// 交换节点
		prev.Next = second
		second.Next = first
		first.Next = next

		// 移动节点
		prev = first
		first = next
		if first == nil {
			break
		}
		second = first.Next
	}

	return dummyNode.Next
}
