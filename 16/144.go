package _16

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历：根左右
// 递归
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	var preOrder func(*TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)

	return res
}

// 迭代解法
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res, stack := make([]int, 0), make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}

		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}

	}

	return res
}
