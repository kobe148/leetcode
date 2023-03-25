package _15

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	smallHead, largeHead := &ListNode{Val: -1}, &ListNode{Val: -1}
	small, large := smallHead, largeHead

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}

		head = head.Next
	}

	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
