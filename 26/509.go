package _26

func fib(n int) int {
	if n <= 1 {
		return n
	}

	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n]
}

// 递归
func fib2(n int) int {
	var dfs func(int) int
	dfs = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		left := dfs(n - 1)
		right := dfs(n - 2)

		return left + right
	}

	return dfs(n)
}

// DFS + 记忆化搜索
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func fib3(n int) int {
	var memo = make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		if memo[n] != -1 {
			return memo[n]
		}
		left := dfs(n - 1)
		right := dfs(n - 2)

		memo[n] = left + right
		return memo[n]
	}

	return dfs(n)
}
