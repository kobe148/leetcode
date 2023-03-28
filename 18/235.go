package _18

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	var left = lowestCommonAncestor(root.Left, p, q)
	var right = lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}
