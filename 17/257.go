package _17

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	res := make([][]int, 0)
	var dfs func(node *TreeNode, path []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
		path = path[:len(path)-1]
	}

	dfs(root, make([]int, 0))

	ret := make([]string, 0, len(res))
	for _, re := range res {
		tmp := strconv.Itoa(re[0])
		for i := 1; i < len(re); i++ {
			tmp += "->" + strconv.Itoa(re[i])
		}
		ret = append(ret, tmp)
	}

	return ret
}

// 前序遍历
func binaryTreePaths2(root *TreeNode) []string {
	res := make([]string, 0)

	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
		}

		dfs(node.Left, path+strconv.Itoa(node.Val)+"->")
		dfs(node.Right, path+strconv.Itoa(node.Val)+"->")
	}

	dfs(root, "")
	return res
}
