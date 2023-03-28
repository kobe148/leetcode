package _19

func maxAreaOfIsland(grid [][]int) int {
	var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	rows, cols := len(grid), len(grid[0])
	var visited = make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	max := 0
	// 后序遍历
	var dfs func(int, int) int
	dfs = func(row int, col int) int {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == 0 || visited[row][col] {
			return 0
		}

		visited[row][col] = true
		ret := 0
		for _, dir := range dirs {
			var nextRow = row + dir[0]
			var nextCol = col + dir[1]
			ret += dfs(nextRow, nextCol)
		}

		return ret + 1
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 && !visited[row][col] {
				res := dfs(row, col)
				if res > max {
					max = res
				}
			}
		}
	}

	return max
}
