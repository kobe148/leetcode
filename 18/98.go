package _18

// 1. 中序遍历(先拿到遍历结果，再验证顺序性)
func isValidBST(root *TreeNode) bool {
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

	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}
