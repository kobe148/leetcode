package _15

func sortList(head *ListNode) *ListNode {
	// 归并排序, 自低朝上,迭代方式：两两合并排序
	if head == nil || head.Next == nil {
		return head
	}

	// 这个需要在计算长度之前处理，因为计算长度之后head就变成了nil了
	var dummy = &ListNode{Next: head}
	// 计算链表长度
	length := 0
	for head != nil {
		head = head.Next
		length++
	}

	// 每次合并的长度
	for step := 1; step < length; step += step {
		// prev用来合并，curr用来切割
		prev, curr := dummy, dummy.Next
		for curr != nil {
			left := curr
			right := split(left, step)
			// 分割得到下次处理的链表头
			curr = split(right, step)
			// 合并 left 和 right 链表
			prev = merge(left, right, prev)
		}
	}

	return dummy.Next
}

// 将 node 从第 step 个节点切断，并返回右边链表的头节点
func split(node *ListNode, step int) *ListNode {
	if node == nil {
		return nil
	}
	// 找到分割点
	for i := 1; node.Next != nil && i < step; i++ {
		node = node.Next
	}
	var right = node.Next
	node.Next = nil

	return right
}

// 合并 left 和 right 两个有序链表
// 将 prev.next 设置为合并之后链表的表头
// 返回合并之后链表的最后一个节点
func merge(left *ListNode, right *ListNode, prev *ListNode) *ListNode {
	var curr = prev
	for left != nil && right != nil {
		if left.Val <= right.Val {
			curr.Next = left
			left = left.Next
		} else {
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}
	if left == nil {
		curr.Next = right
	}
	if right == nil {
		curr.Next = left
	}

	// 保证 curr 是合并之后链表的最后一个节点
	for curr.Next != nil {
		curr = curr.Next
	}

	return curr
}
