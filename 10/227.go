package _10

func calculate227(s string) int {
	stack := make([]int, 0, len(s))

	preSign, i := byte('+'), 0

	for i < len(s) {
		if s[i] == ' ' {
			i++
		} else {
			// 获取数字
			num := 0
			for i < len(s) && (s[i] >= '0' && s[i] <= '9') {
				num = num*10 + int(s[i]-'0')
				i++
			}

			if preSign == '+' {
				stack = append(stack, num)
			} else if preSign == '-' {
				stack = append(stack, -num)
			} else if preSign == '*' {
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, tmp*num)
			} else if preSign == '/' {
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, tmp/num)
			}

			for i < len(s) && s[i] == ' ' {
				i++
			}

			if i < len(s) {
				preSign = s[i]
			}

			i++
		}

	}

	res := 0
	for _, num := range stack {
		res += num
	}

	return res
}
