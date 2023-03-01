package _08

func searchMatrix(matrix [][]int, target int) bool {
	// 二维数组变成一维数组来二分查找
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		num := matrix[mid/n][mid%n]
		if num == target {
			return true
		} else if target > num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
