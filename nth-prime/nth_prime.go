package prime

// Nth returns the nth prime number
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	current := 1
	c := make(chan int)
	go getPrimes(c)

	for prime := range c {
		if current == n {
			return prime, true
		}
		current++
	}

	// will never reach here
	return 0, false
}

func getPrimes(c chan int) {
	limit := 1000000
	for i := 0; i < limit; i++ {
		func(i int) {
			if isPrime(i) {
				c <- i
			}
		}(i)
	}
}

func isPrime(n int) bool {
	// Corner cases
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	// This is checked so that we can skip
	// middle five numbers in below loop
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
