package sieve

// 4319 ns/op

// Sieve of Eratosthenes, a method for finding primes up to a limit not using div, /, mod or %
func Sieve(limit int) (primes []int) {
	nonPrime := make([]bool, limit+1)
	if limit > 1 {
		// append 2 so we can avoid checking multiples of it below
		primes = append(primes, 2)
	}

	for i := 3; i <= limit; i += 2 {
		if !nonPrime[i] {
			primes = append(primes, i)
			appendSelfAndMultiples(i, limit, &nonPrime)
		}
	}

	return primes
}

func appendSelfAndMultiples(num int, limit int, collection *[]bool) {
	// we only care about from the next square value,
	// e.g. for num = 6, no point in adding 6 itself since we're past it, plus anything below 36 will have
	// been added already on the lower side of the multiple (e.g. 12 will have been added by 2 x 6, 18 by 3 x 6)
	current := num * num
	for current <= limit {
		(*collection)[current] = true
		current = current + num
	}
}
