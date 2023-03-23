package _14

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := head.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next // 移动

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}
