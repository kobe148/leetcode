package _14

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, curr := head, head.Next

	for curr != nil {
		if prev.Val == curr.Val {
			prev.Next = curr.Next
			curr.Next = nil
		} else {
			prev = curr
		}

		curr = prev.Next
	}

	return head
}
