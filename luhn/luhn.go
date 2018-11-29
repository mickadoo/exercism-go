package luhn

import (
	"errors"
	"strings"
	"unicode"
)

// Valid checks if the given string is a valid Luhn number
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)

	if len(input) < 2 {
		return false
	}

	numbers, err := getIntSlice(input)

	if err != nil {
		return false
	}

	change := false
	for i := len(numbers) - 1; i >= 0; i-- {
		if change == true {
			numbers[i] = (numbers[i] * 2)
			if numbers[i] > 9 {
				numbers[i] -= 9
			}
			change = false
		} else {
			change = true
		}
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum%10 == 0
}

func getIntSlice(input string) (result []int, err error) {
	for _, char := range input {
		if unicode.IsDigit(char) {
			result = append(result, int(char-'0'))
		} else if unicode.IsSpace(char) {
			continue
		} else {
			err = errors.New("Contains non digit character")
			return
		}
	}

	return
}
