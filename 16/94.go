package _16

// 递归
func inorderTraversal(root *TreeNode) []int {
	var res = make([]int, 0)

	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}

	inOrder(root)
	return res

}

// 迭代:左根右
func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	curr := root

	for curr != nil || len(stack) > 0 {
		// 先找到最左边的节点
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// 处理栈顶元素
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		curr = node.Right
	}

	return res
}
