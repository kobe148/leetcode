package _19

func postorder(root *Node) []int {
	var res = make([]int, 0)

	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			dfs(child)
		}
		res = append(res, node.Val)
	}

	dfs(root)

	return res
}
