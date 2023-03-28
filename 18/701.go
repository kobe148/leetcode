package _18

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	curr := root
	for curr != nil {
		if val < curr.Val && curr.Left == nil {
			curr.Left = &TreeNode{Val: val}
			return root
		} else if val > curr.Val && curr.Right == nil {
			curr.Right = &TreeNode{Val: val}
			return root
		}

		if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return root
}
