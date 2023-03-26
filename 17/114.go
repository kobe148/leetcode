package _17

func flatten(root *TreeNode) {
	res := make([]*TreeNode, 0)

	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node)
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)

	// 串联成链表
	for i := 1; i < len(res); i++ {
		prev, curr := res[i-1], res[i]
		prev.Left = nil
		prev.Right = curr
	}
}
