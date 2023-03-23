package _14

func deleteDuplicates82(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: -1, Next: head}

	prev, curr := dummy, head
	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			prev.Next = curr.Next
			curr.Next = nil
		} else {
			prev = curr
		}

		curr = prev.Next
	}

	return dummy.Next
}
