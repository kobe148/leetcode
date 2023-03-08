package _10

func calculate(s string) int {
	stack := make([]int, 0, len(s))

	num := 0
	res := 0
	preSign := 1

	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == '+' {
			res += preSign * num
			preSign, num = 1, 0
		} else if c == '-' {
			res += preSign * num
			preSign, num = -1, 0
		} else if c == '(' {
			stack = append(stack, res)
			stack = append(stack, preSign)
			preSign, res = 1, 0
		} else if c == ')' {
			res += preSign * num
			res *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num = 0
		}
	}

	return res + preSign*num
}
