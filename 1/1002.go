package _01

func CommonChars(words []string) []string {
	res := make([]string, 0, len(words))
	chars := [26]int{}
	// 最少次数
	// 先统计第一个字符
	for i := 0; i < len(words[0]); i++ {
		chars[words[0][i]-'a']++
	}
	for i := 1; i < len(words); i++ {
		min := [26]int{}
		for _, c := range words[i] {
			min[c-'a']++
		}
		for i := 0; i < len(min); i++ {
			if min[i] < chars[i] {
				chars[i] = min[i]
			}
		}
	}

	for i := 0; i < len(chars); i++ {
		for j := 0; j < chars[i]; j++ {
			res = append(res, string(rune(i+'a')))
		}
	}
	return res
}
