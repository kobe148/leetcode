package _03

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for row := 0; row < numRows; row++ {
		rows := make([]int, 0)
		for col := 0; col <= row; col++ {
			if col == 0 || col == row {
				rows = append(rows, 1)
			} else {
				preRow := res[row-1]
				rows = append(rows, preRow[col-1]+preRow[col])
			}
		}
		res = append(res, rows)
	}

	return res
}
