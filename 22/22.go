package _22

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n <= 0 {
		return res
	}

	var dfs func(string, int, int)
	dfs = func(path string, open int, close int) {
		if len(path) == 2*n {
			res = append(res, path)
		}
		// 剪枝策略
		if open < n {
			dfs(path+"(", open+1, close)
		}
		if close < open {
			dfs(path+")", open, close+1)
		}

	}

	dfs("", 0, 0)
	return res
}
