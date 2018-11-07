// Package romannumerals deals with translating arabic numbers to roman numerals
package romannumerals

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// numeralsByLevel has list of roman numerals (for 1, 5 and 10) for each the nth digit levels
var numeralsByLevel = [][]string{
	{"I", "V", "X"}, // for last digit (singles), e.g. "4" in 5134
	{"X", "L", "C"}, // for second last digit (tens), e.g. "3" in 5134
	{"C", "D", "M"}, // for third last digit (hundreds), e.g. "1" in 5134
	{"M", "ↁ", "ↂ"}, // for fourth last digit (thousands), e.g. "5" in 5134
}

// ToRomanNumeral takes an number and returns the roman numeral equivalent
func ToRomanNumeral(num int) (string, error) {
	if num < 1 || num > 3000 {
		return "", errors.New("Number must be between 1 and 300")
	}

	result := ""
	for n := range numeralsByLevel {
		nthDigit := getNthDigit(num, n)
		result = digitToRoman(nthDigit, n) + result
	}

	return result, nil
}

// The nth digit is (the remainder of dividing by 10n) divided by 10n-1
func getNthDigit(num, n int) int {
	return num / int(math.Pow10(n)) % 10
}

// digitToRoman returns the roman representation for a single decimal digit based on its position in the number
// For example 2 will be "II" if it is the last digit in the number, but "XX" if n = 1
func digitToRoman(digit int, n int) string {
	numerals := numeralsByLevel[n]
	charForOne := numerals[0]
	charForFive := numerals[1]
	charForTen := numerals[2]

	switch {
	case digit < 4:
		return strings.Repeat(charForOne, digit)
	case digit == 4:
		return charForOne + charForFive
	case digit == 5:
		return charForFive
	case digit < 9:
		return charForFive + strings.Repeat(charForOne, digit%5)
	case digit == 9:
		return charForOne + charForTen
	case digit == 0:
		return ""
	}

	panic(fmt.Sprintf("Could not handle this digit - %d", digit))
}
