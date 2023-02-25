package _07

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	return quickPow(x, n)
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	mid := n / 2

	y := quickPow(x, mid)

	if n%2 == 0 {
		return y * y
	}

	return x * y * y
}
