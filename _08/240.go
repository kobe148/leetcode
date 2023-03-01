package _08

func searchMatrix240(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	shortDim := min(m, n)

	for i := 0; i < shortDim; i++ {
		rowFound := binarySearchRow(matrix, i, target)
		colFound := binarySearchCol(matrix, i, target)

		if rowFound || colFound {
			return true
		}
	}

	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func binarySearchRow(matrix [][]int, row int, target int) bool {
	var left, right = row, len(matrix[0]) - 1
	for left <= right {
		var mid = left + (right-left)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func binarySearchCol(matrix [][]int, col int, target int) bool {
	var left, right = col, len(matrix) - 1
	for left <= right {
		var mid = left + (right-left)/2
		if matrix[mid][col] == target {
			return true
		} else if matrix[mid][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
