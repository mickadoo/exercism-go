package luhn

import (
	"strings"
	"unicode"
)

// 203 ns/op

// Valid checks if the given string is a valid Luhn number
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)

	if len(input) < 2 {
		return false
	}

	flag := 1
	if len(input)%2 == 0 {
		flag = 0
	}

	var digit int
	sum := 0
	for index, char := range input {
		if unicode.IsDigit(char) {
			digit = int(char - '0')
			if index%2 == flag {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
		} else if unicode.IsSpace(char) {
			continue
		} else {
			return false
		}
	}

	return sum%10 == 0
}
