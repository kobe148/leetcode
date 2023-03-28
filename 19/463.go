package _19

// DFS
func islandPerimeter(grid [][]int) int {
	var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	rows, cols := len(grid), len(grid[0])
	var visited = make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// 后序遍历
	var dfs func(int, int) int
	dfs = func(row int, col int) int {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == 0 {
			return 1
		}

		if visited[row][col] {
			return 0
		}

		visited[row][col] = true
		res := 0
		for _, dir := range dirs {
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			res += dfs(nextRow, nextCol)
		}

		return res
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				return dfs(row, col)
			}
		}
	}

	return 0
}
