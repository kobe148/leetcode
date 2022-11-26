package _03

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	flagCol1 := false

	for row := 0; row < m; row++ {
		if matrix[row][0] == 0 {
			flagCol1 = true
		}
		for col := 1; col < n; col++ {
			if matrix[row][col] == 0 {
				matrix[0][col] = 0
				matrix[row][0] = 0
			}
		}
	}

	for row := m - 1; row >= 0; row-- {
		for col := 1; col < n; col++ {
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}
		}
		if flagCol1 {
			matrix[row][0] = 0
		}
	}
}
