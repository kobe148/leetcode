package _06

func isPowerOfTwo(n int) bool {
	// 使用除法

	// 2的幂，二进制只有一个1
	if n == 0 {
		return true
	}
	return (n & (n - 1)) == 0
}
