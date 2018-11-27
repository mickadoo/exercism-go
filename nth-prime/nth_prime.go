package prime

// Nth returns the nth prime number
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	current := 0
	for i := 2; true; i++ {
		if isPrime(i) {
			current++
			if current == n {
				return i, true
			}
		}
	}

	// will never reach here
	return 0, false
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
