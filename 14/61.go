package _14

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 计算链表的长度
	var length, curr = 0, head
	for curr != nil {
		length++
		curr = curr.Next
	}

	if k%length == 0 {
		return head
	}

	k = k % length
	var newHead = reverseList(head)
	var kthNode = newHead
	for i := 0; i < k-1; i++ {
		kthNode = kthNode.Next
	}

	var nextHead = kthNode.Next
	kthNode.Next = nil

	var retHead = reverseList(newHead)
	newHead.Next = reverseList(nextHead)

	return retHead
}
