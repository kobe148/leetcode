package _19

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := make([]int, 0)

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for i := 0; i < len(node.Children); i++ {
			dfs(node.Children[i])
		}
	}

	dfs(root)
	return res
}
