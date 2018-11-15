// Package anagram is used for detecting anagrams
package anagram

import "strings"

// Detect checks for anagrams of a given word in a list of candidates
func Detect(subject string, candidates []string) (result []string) {
	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			result = append(result, candidate)
		}
	}
	return
}

func isAnagram(original string, anagram string) bool {
	// make them both the same case
	original = strings.ToLower(original)
	anagram = strings.ToLower(anagram)

	// words cannot be anagrams of themselves
	if original == anagram {
		return false
	}

	for _, char := range original {
		index := strings.IndexRune(anagram, char)
		if index == -1 {
			return false
		}
		// remove matching character as it can only match once
		anagram = anagram[0:index] + anagram[index+1:]
	}

	return anagram == ""
}
