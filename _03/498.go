package _03

func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 {
		return []int{}
	}

	m, n := len(mat), len(mat[0])
	dirs := [][]int{{-1, 1}, {1, -1}}

	res := make([]int, 0, m*n)
	row, col, di := 0, 0, 0
	for i := 0; i < m*n; i++ {
		res = append(res, mat[row][col])

		row = row + dirs[di][0]
		col = col + dirs[di][1]

		if col >= n {
			col = n - 1
			row += 2
			di = 1 - di
		}
		if row >= m {
			row = m - 1
			col += 2
			di = 1 - di
		}
		if row < 0 {
			row = 0
			di = 1 - di
		}
		if col < 0 {
			col = 0
			di = 1 - di
		}
	}

	return res
}
