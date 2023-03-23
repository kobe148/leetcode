package _14

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 || k == 1 {
		return head
	}

	dummyNode := &ListNode{Next: head}

	prev, first, last := dummyNode, head, head

	for first != nil {
		for i := 0; i < k-1; i++ {
			last = last.Next
			if last == nil {
				return dummyNode.Next
			}
		}

		next := last.Next
		last.Next = nil

		reverseList(first)

		prev.Next, first.Next = last, next

		prev = first
		first = next
		last = first
	}

	return dummyNode.Next
}
