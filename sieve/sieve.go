package sieve

// 268539 ns/op

// Sieve of Eratosthenes, a method for finding primes up to a limit not using div, /, mod or %
func Sieve(limit int) (primes []int) {
	var nonPrime []int
	for i := 2; i <= limit; i++ {
		if !exists(i, nonPrime) {
			primes = append(primes, i)
			appendSelfAndMultiples(i, limit, &nonPrime)
		}
	}

	return primes
}

func exists(num int, collection []int) bool {
	for _, val := range collection {
		if val == num {
			return true
		}
	}

	return false
}

func appendSelfAndMultiples(num int, limit int, collection *[]int) {
	current := num
	i := 1
	for current <= limit {
		*collection = append(*collection, current)
		current = num * i
		i++
	}
}
