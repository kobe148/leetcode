package _01

func SortString(s string) string {
	sArr := [26]int{}
	// 统计出每个字符有多少个
	for _, c := range s {
		sArr[c-'a']++
	}
	var res string
	// 先从小到大排序
	for len(res) < len(s) {
		for i := 0; i < len(sArr); i++ {
			if sArr[i] > 0 {
				res += string(rune(i + 'a'))
				sArr[i]--
			}
		}
		for i := len(sArr) - 1; i >= 0; i-- {
			if sArr[i] > 0 {
				res += string(rune(i + 'a'))
				sArr[i]--
			}
		}
	}

	return res
}
