package _14

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	// fast先走N + 1步
	for n >= 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	delNode := slow.Next
	slow.Next = delNode.Next
	delNode.Next = nil

	return dummy.Next
}
