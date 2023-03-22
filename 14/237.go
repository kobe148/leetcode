package _14

func deleteNode(node *ListNode) {
	prev := &ListNode{}

	for node != nil {
		next := node.Next
		if next != nil {
			node.Val = next.Val
		} else {
			prev.Next = nil
		}

		prev = node
		node = node.Next
	}
}
