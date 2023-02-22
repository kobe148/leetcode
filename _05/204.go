package _05

func countPrimes(n int) int {
	notPrimes := make([]bool, n)
	ans := 0
	for i := 2; i < n; i++ {
		if notPrimes[i] {
			continue
		}
		ans++

		for j := i; j < n; j += i {
			notPrimes[j] = true
		}
	}

	return ans
}
