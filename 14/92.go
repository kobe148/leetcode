package _14

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	dummy := &ListNode{Next: head}

	// 1. 从虚拟节点走 left - 1 步，来到 left 节点的前一个节点
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	leftNode := prev.Next

	// 2. 从 leftNode 节点开始走 right - left 步，来到 right 节点
	rightNode := leftNode
	for i := 0; i < right-left; i++ {
		rightNode = rightNode.Next
	}

	postRight := rightNode.Next

	// 3. 切断得到 left 到 right 的子链表
	prev.Next = nil
	rightNode.Next = nil

	// 4. 反转 leftNode 到 rightNode
	reverseList(leftNode)

	// 5. 将反转后的子链表接回到原链表
	prev.Next = rightNode
	leftNode.Next = postRight

	return dummy.Next
}
