package _15

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l1 = reverseList(l1)
	l2 = reverseList(l2)

	var retNode = addTwoNumbersHelp(l1, l2)
	return reverseList(retNode)
}

func addTwoNumbersHelp(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	curr, carry := dummy, 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry
		curr.Next = &ListNode{Val: sum % 10}
		// bug 修复：视频中忘了加上这一步
		curr = curr.Next
		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
