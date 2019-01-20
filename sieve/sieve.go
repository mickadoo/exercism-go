package sieve

// 5651 ns/op

// Sieve of Eratosthenes, a method for finding primes up to a limit not using div, /, mod or %
func Sieve(limit int) (primes []int) {
	nonPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		if !nonPrime[i] {
			primes = append(primes, i)
			appendSelfAndMultiples(i, limit, &nonPrime)
		}
	}

	return primes
}

func appendSelfAndMultiples(num int, limit int, collection *[]bool) {
	current := num
	for current <= limit {
		(*collection)[current] = true
		current = current + num
	}
}
