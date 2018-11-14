// Package pangram is involved with checking pangrams
package pangram

import (
	"strings"
)

// IsPangram checks if the word provided is a pangram, i.e. it contains all letters from a-z
func IsPangram(str string) bool {
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	str = strings.ToLower(str)

	for _, letter := range letters {
		if !strings.ContainsRune(str, letter) {
			return false
		}
	}

	return true
}
