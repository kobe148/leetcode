package _04

func convert(s string, numRows int) string {
	temp := make([]string, numRows)

	i, n := 0, len(s)
	for i < n {
		// 垂直向下
		for idx := 0; idx < numRows && i < n; idx++ {
			temp[idx] += string(s[i])
			i++
		}
		// 右向上
		for idx := numRows - 2; idx >= 1 && i < n; idx-- {
			temp[idx] += string(s[i])
			i++
		}
	}

	res := ""
	for j := 0; j < len(temp); j++ {
		res += temp[j]
	}

	return res
}
