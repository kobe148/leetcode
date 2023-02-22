package _05

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	head := dummy
	carry := 0

	for l1.Next != nil || l2.Next != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry
		carry = sum / 10

		head.Next = &ListNode{Val: sum % 10}
		head = head.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry > 0 {
		head.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
