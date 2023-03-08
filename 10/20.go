package _10

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			// 进来这里的都是 ')',  ']', '}'， 这样的是不需要入栈的，需要将栈顶元素弹出来匹配
			// 如果栈的长度 == 0 ，表示没有数据可以匹配了
			if len(stack) == 0 {
				return false
			}
			// 出栈匹配
			top := stack[len(stack)-1]
			// 查看是否匹配
			stack = stack[:len(stack)-1]

			if s[i] == ')' && top != '(' {
				return false
			}
			if s[i] == '}' && top != '{' {
				return false
			}
			if s[i] == ']' && top != '[' {
				return false
			}
		}
	}

	// 最终栈的长度看是否为0，是的话则表示完全匹配
	return len(stack) == 0
}
