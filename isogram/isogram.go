// Package isogram deals with isogram words
package isogram

import (
	"unicode"
)

// IsIsogram checks if a word is an isogram, i.e. it has no repeating letters. It only works for words with ascii characters
func IsIsogram(word string) bool {
	var charactersFound int32

	for _, letter := range word {
		letter = unicode.ToLower(letter)

		// get the numeric representation (0-25) by subtracting the lowest alphabetic character
		flag := uint(letter - 'a')

		// ignore characters outside of a-z
		if flag > 25 {
			continue
		}

		// check if this flag was already set
		if charactersFound&(1<<flag) > 0 {
			return false
		}
		// set the bit for this flag to 1
		charactersFound ^= 1 << flag
	}

	return true
}
