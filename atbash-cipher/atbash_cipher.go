package atbash

import (
	"strings"
)

// Atbash encodes the input using the atbash cipher
func Atbash(input string) (ouput string) {
	input = strings.ToLower(input)
	code := []rune{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}
	spaceCount := 0
	for _, char := range input {
		index := char - 'a'
		if index >= 0 && index <= 26 {
			ouput += string(code[index])
			if (len(ouput)-spaceCount)%5 == 0 {
				ouput += " "
				spaceCount++
			}
		}
		numIndex := char - '0'
		if numIndex >= 0 && numIndex <= 9 {
			ouput += string(char)
			if (len(ouput)-spaceCount)%5 == 0 {
				ouput += " "
				spaceCount++
			}
		}
	}

	return strings.Trim(ouput, " ")
}
