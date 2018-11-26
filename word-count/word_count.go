package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is a map of words to their frequency counts
type Frequency map[string]int

// WordCount counts the number of words in the given input
func WordCount(input string) Frequency {
	input = strings.ToLower(input)

	// remove quotations
	input = strings.Replace(input, " '", " ", -1)
	input = strings.Replace(input, "' ", " ", -1)

	words := strings.FieldsFunc(input, func(r rune) bool {
		// don't split on apostrophe
		if r == '\'' {
			return false
		}
		return unicode.IsSpace(r) || unicode.IsSymbol(r) || unicode.IsPunct(r)
	})

	result := make(Frequency)
	for _, word := range words {
		result[word]++
	}

	return result
}
