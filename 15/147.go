package _15

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &ListNode{Next: head}
	// curr是当前需要处理的节点，prev是已经处理的节点
	prev, curr := head, head.Next

	// head = [4,2,1,3]
	// 输出: [1,2,3,4]
	for curr != nil {
		if curr.Val >= prev.Val {
			prev = curr
			curr = curr.Next
		} else {
			// 说明：这里的 p.next 不可能为空
			// 因为 p 从头开始，最远可以到达的节点是 curr 的前一个节点
			// 所以 p.next 不可能为 nil，我这里加上 p.next 的判空，是我个人的习惯哟~
			var p = dummyNode

			// 找到小于 curr.val 的最大的节点
			for p.Next != nil && p.Next.Val < curr.Val {
				p = p.Next
			}

			// 记录下一个要处理的元素以及下一个节点
			prev.Next = curr.Next
			// 将 curr 插入到 p 和 p.next 之间
			curr.Next = p.Next
			p.Next = curr
			// 下一个需要比较的节点,因为curr经过交换后的位置有变化，其实下一个要要处理的元素就是prev.Next
			curr = prev.Next
		}

	}

	return dummyNode.Next
}
