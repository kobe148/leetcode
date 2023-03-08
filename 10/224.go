package _10

func calculate(s string) int {
	numStack := make([]int, 0, len(s))

	num := 0
	res := 0
	preSign := 1

	for _, c := range s {
		if c >= '0' && c <= '9' {
			// 数字
			num = num*10 + int(c-'0')
		} else if c == '+' {
			res += preSign * num
			preSign, num = 1, 0
		} else if c == '-' {
			res += preSign * num
			preSign, num = -1, 0
		} else if c == '(' {
			numStack = append(numStack, res)
			numStack = append(numStack, preSign)
			preSign, num = 1, 0
		} else if c == ')' {
			res += preSign * num
			res *= numStack[len(numStack)-1] // 这个是符号
			numStack = numStack[:len(numStack)-1]

			res += numStack[len(numStack)-1] // 这个是数字
			numStack = numStack[:len(numStack)-1]

			num = 0
		}
	}

	return res + preSign*num
}
