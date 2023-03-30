package _22

func letterCombinations(digits string) []string {
	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}

	var dfs func(int, string)
	dfs = func(index int, combination string) {
		// index是指digits字符串的索引, 比如digits = "23"
		// index == 0, 表示numChar = 2， letters = abc
		if len(digits) == index {
			res = append(res, combination)
			return
		}

		numChar := digits[index]
		letters := phone[numChar]
		for i := range letters {
			dfs(index+1, combination+string(letters[i]))
		}
	}

	dfs(0, "")

	return res
}
