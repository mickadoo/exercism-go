// Package summultiples is used for calculating the sum of multiples
package summultiples

import "fmt"

/*
SumMultiples finds the sum of all the unique multiples of particular
numbers up to but not including that number.
*/
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	common := make([]int, 0)

	for _, outer := range divisors {
		for _, inner := range divisors {
			if outer == inner {
				continue
			}
			new := lcm(outer, inner)
			if !exists(common, new) {
				common = append(common, new)
			}
		}
	}

	fmt.Println(common)

	for _, divisor := range divisors {
		sum += formula(divisor, limit-1)
	}

	for _, commonVal := range common {
		sum -= formula(commonVal, limit-1)
	}

	return sum
}

func formula(n int, limit int) int {
	return n * int(limit/n) * (1 + int(limit/n)) / 2
}

func exists(slice []int, val int) bool {
	for _, existing := range slice {
		if existing == val {
			return true
		}
	}
	return false
}

// Recursive function to
// return gcd of a and b
func gcd(a, b int) int {
	// base case
	if a == b {
		return a
	}

	// a is greater
	if a > b {
		return gcd(a-b, b)
	}
	return gcd(a, b-a)
}

// Function to return LCM
// of two numbers
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
