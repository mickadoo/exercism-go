// Package isogram deals with isogram words
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if a word is an isogram, i.e. it has no repeating letters
func IsIsogram(word string) bool {
	charCounts := make(map[rune]int)
	word = strings.ToLower(word)

	for _, letter := range word {
		if unicode.IsLetter(letter) {
			charCounts[letter]++
		}
		if charCounts[letter] > 1 {
			return false
		}
	}

	return true
}
