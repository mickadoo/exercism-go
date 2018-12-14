package atbash

import (
	"strings"
	"unicode"
)

// 7315 ns/op

// Atbash encodes the input using the atbash cipher
func Atbash(input string) string {
	input = strings.ToLower(input)
	output := make([]rune, 0)
	code := []rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}
	for _, char := range input {
		index := char - 'a'
		if index >= 0 && index <= 26 {
			output = append(output, code[index])
		} else if unicode.IsNumber(char) {
			output = append(output, char)
		}
		if (len(output))%6 == 5 {
			output = append(output, ' ')
		}
	}

	return strings.Trim(string(output), " ")
}
