// Package summultiples is used for calculating the sum of multiples
package summultiples

/*
SumMultiples finds the sum of all the unique multiples of particular
numbers up to but not including that number.
*/
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if i%divisor == 0 {
				sum += i
				break // don't count same num twice
			}
		}
	}
	return sum
}
