package atbash

import (
	"bytes"
	"strings"
	"unicode"
)

// Atbash encodes the input using the atbash cipher
func Atbash(input string) string {
	input = strings.ToLower(input)
	code := []rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}
	var buffer bytes.Buffer

	for _, char := range input {
		index := char - 'a'
		if index >= 0 && index <= 26 {
			buffer.WriteRune(code[index])
		} else if unicode.IsNumber(char) {
			buffer.WriteRune(char)
		}
		if (buffer.Len())%6 == 5 {
			buffer.WriteRune(' ')
		}
	}

	return strings.Trim(buffer.String(), " ")
}
