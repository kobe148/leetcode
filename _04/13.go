package _04

func romanToInt(s string) int {
	symbolValues := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	// 比较前一个字符跟当前字符那个大
	sum := 0
	pre := symbolValues[s[0]]

	for i := 1; i < len(s); i++ {
		curr := symbolValues[s[i]]
		if pre < curr {
			sum -= pre
		} else {
			sum += pre
		}
		pre = curr // 字符串要更新位置

	}

	sum += pre // 最后一个位置的数字

	return sum
}
