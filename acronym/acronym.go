// Package acronym generates acronyms
package acronym

import "strings"

// Abbreviate takes a string and returns its acronym
func Abbreviate(fullString string) string {
	fullString = strings.Replace(fullString, "-", " ", -1)
	words := strings.Fields(fullString)

	acronym := ""
	for _, word := range words {
		acronym += string(word[0])
	}

	return strings.ToUpper(acronym)
}
