package _18

// 中序遍历
// 另外一种解法：直接递归树把接节点存入数组，直接在数组中计算
func findMode(root *TreeNode) []int {
	ans := make([]int, 0)
	prevNum, count, maxCount := 0, 0, 0

	update := func(val int) {
		if val == prevNum {
			count++
		} else {
			prevNum = val
			count = 1
		}

		if count == maxCount {
			ans = append(ans, val)
		} else if count > maxCount {
			ans = make([]int, 0)
			ans = append(ans, val)
			maxCount = count
		}
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}

	dfs(root)

	return ans
}
