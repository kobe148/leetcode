package _03

// 输入：matrix =
// [
//
//		[1,2,3],
//		[4,5,6],
//		[7,8,9]
//	]
//
// 输出：[
//
//			   [7,4,1],
//			   [8,5,2],
//			   [9,6,3]
//			]
//		 转换前： 2 ==> matrix[0, 1]    转换后：2 ==> matrix[1, 2]
//		 转换前： 6 ==> matrix[1, 2]    转换后：6 ==> matrix[2, 1]
//		 转换前： 4 ==> matrix[1, 0]    转换后：4 ==> matrix[0, 1]
//		 转换前： 8 ==> matrix[2, 1]    转换后：8 ==> matrix[1, 0]
//	     转换公式: matrix[row][col] -> matrix[col][n - row - 1]
func rotate(matrix [][]int) {
	n := len(matrix)

	// 每一轮转换都是用上面的公式把row，col代入即可
	for row := 0; row < n/2; row++ {
		for col := 0; col < (n+1)/2; col++ {
			tmp := matrix[row][col]
			matrix[row][col] = matrix[n-col-1][row]
			matrix[n-col-1][row] = matrix[n-row-1][n-col-1]
			matrix[n-row-1][n-col-1] = matrix[col][n-row-1]
			matrix[col][n-row-1] = tmp
		}
	}
}
