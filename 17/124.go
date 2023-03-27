package _17

import "math"

func maxPathSum(root *TreeNode) int {
	// 自底向上
	var maxSum = math.MinInt32

	// 返回以 node 为根节点的子树的最大贡献值
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		left := maxGain(node.Left)
		if left < 0 {
			left = 0
		}
		right := maxGain(node.Right)
		if right < 0 {
			right = 0
		}

		if left+right+node.Val > maxSum {
			maxSum = left + right + node.Val
		}

		if left > right {
			return left + node.Val
		} else {
			return right + node.Val
		}
	}

	maxGain(root)

	return maxSum
}
