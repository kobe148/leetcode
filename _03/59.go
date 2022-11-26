package _03

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
		res[i] = make([]int, n)
	}
	// 控制方向
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	di := 0
	row, col := 0, 0
	for i := 0; i < n*n; i++ {
		res[row][col] = i + 1
		visited[row][col] = true
		// 判断
		nextRow := row + dirs[di][0]
		nextCol := col + dirs[di][1]
		// 超出范围
		if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			di = (di + 1) % 4
		}

		row = row + dirs[di][0]
		col = col + dirs[di][1]

	}

	return res
}
