// Package diffsquares is used to calculate the diff of squared numbers
package diffsquares

/*
SquareOfSum is the square of the sum of the first n natural numbers
*/
func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

/*
SumOfSquares is the sum of the squares of the first n natural numbers
*/
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

/*
Difference is the difference between the square of the sum of the first
n natural numbers and the sum of the squares of the first n
natural numbers
*/
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
