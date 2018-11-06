// Package raindrops is involved in converting numbers to pattern-based strings
package raindrops

import (
	"strconv"
)

// Convert takes a number and converts it to a modified string based on its factors
func Convert(num int) string {
	result := ""

	if num%3 == 0 {
		result = "Pling"
	}

	if num%5 == 0 {
		result += "Plang"
	}

	if num%7 == 0 {
		result += "Plong"
	}

	if result != "" {
		return result
	}

	return strconv.Itoa(num)
}
