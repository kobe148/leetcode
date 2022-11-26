package _03

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	// 控制方向
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	di := 0
	row, col := 0, 0
	res := make([]int, 0, m*n)
	for i := 0; i < m*n; i++ {
		res = append(res, matrix[row][col])
		visited[row][col] = true
		// 判断
		nextRow := row + dirs[di][0]
		nextCol := col + dirs[di][1]
		// 超出范围
		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			di = (di + 1) % 4
		}

		row = row + dirs[di][0]
		col = col + dirs[di][1]

	}
	return res
}
