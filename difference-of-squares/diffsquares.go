// Package diffsquares is used to calculate the diff of squared numbers
package diffsquares

/*
SquareOfSum is the square of the sum of the first n natural numbers
*/
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

/*
SumOfSquares is the sum of the squares of the first n natural numbers
*/
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

/*
Difference is the difference between the square of the sum of the first
n natural numbers and the sum of the squares of the first n
natural numbers
*/
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
