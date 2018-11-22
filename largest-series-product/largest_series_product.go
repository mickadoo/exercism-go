// Package lsproduct is used for calculating products
package lsproduct

import "errors"

/*
LargestSeriesProduct calculates the largest product for a contiguous
substring of digits of length <span> from a given string of digits.
*/
func LargestSeriesProduct(digitString string, span int) (max int64, err error) {
	digits, err := strToIntSlice(digitString)

	if err != nil {
		return
	}

	if len(digits) < span {
		err = errors.New("Span is greater than string length")
		return
	}

	if span < 0 {
		err = errors.New("Span cannot be negative")
		return
	}

	for i := span; i <= len(digits); i++ {
		current := getProductOfLastN(digits, i, span)
		if current > max {
			max = current
		}
	}

	return
}

func getProductOfLastN(digits []int, offset int, n int) int64 {
	var res int64 = 1
	for n > 0 {
		res *= int64(digits[offset-n])
		n--
	}

	return res
}

func strToIntSlice(digits string) ([]int, error) {
	result := make([]int, len(digits))
	for index, char := range digits {
		intVal := int(char - '0')
		if intVal > 9 || intVal < 0 {
			return result, errors.New("String contains non-numeric char")
		}

		result[index] = intVal
	}

	return result, nil
}
