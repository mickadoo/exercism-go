package luhn

import (
	"strings"
)

// Valid checks if the given string is a valid Luhn number
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)

	if len(input) < 2 {
		return false
	}

	// tracks if digit should be doubled based on input length
	flag := 1
	if len(input)%2 == 0 {
		flag = 0
	}

	var digit int
	sum := 0
	for index, char := range input {
		digit = int(char - '0')

		if digit > 9 || digit < 0 {
			return false
		}

		if index%2 == flag {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
