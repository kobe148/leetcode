package _05

func countDigitOne(n int) int {
	// 暴力：把数字变成string类型，然后在遍历的时候看看字符串类型里面有没有1

	// (n/10)*1 + min(max((n%10 - 1 + 1), 0), 1)
	// (n/100)*10 + min(max((n%100 - 10 + 1), 0), 10)
	// (n/1000)*100 + min(max((n%1000 - 100 + 1), 0), 100)

	count := 0
	for i := 1; i <= n; i *= 10 {
		devider := i * 10
		count += (n/devider)*i + min(max(n%devider-i+1, 0), i)
	}

	return count
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
