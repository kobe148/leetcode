package _16

// 后序：左右根
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		res = append(res, node.Val)
	}

	postOrder(root)

	return res
}

// 迭代解法
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res, stack = make([]int, 0), make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		var curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}

	reverse(res)
	return res
}

func reverse(arr []int) {
	var left, right = 0, len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
