package _10

func decodeString(s string) string {
	stack := make([]int, 0, len(s))
	strStack := make([]string, 0, len(s))

	res := "" // 临时字符串
	num := 0

	for _, char := range s {
		// 数字的情况，直接入栈, 数字可能大于10
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else if char == '[' {
			stack = append(stack, num)
			strStack = append(strStack, res)
			// 这是因为遇到了字符串，所以需要重新统计
			res, num = "", 0
		} else if char == ']' {
			// 处理字符叠加
			// 先把数字出栈
			item := res
			cnt := stack[len(stack)-1]

			for i := 1; i < cnt; i++ {
				// 这时候的currStr是类似2[bc]里面的bc了
				res += item
			}

			// 字符串操作重新入队
			res = strStack[len(strStack)-1] + res
			stack = stack[:len(stack)-1]
			strStack = strStack[:len(strStack)-1]
		} else {
			// 字符的情况
			res += string(char)
		}
	}

	return res
}
