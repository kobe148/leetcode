package _16

// DFS
func isSymmetric(root *TreeNode) bool {
	return isSymmetricR(root, root)
}

func isSymmetricR(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSymmetricR(p.Left, q.Right) && isSymmetricR(p.Right, q.Left)
}
