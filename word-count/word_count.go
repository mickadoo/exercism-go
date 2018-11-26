package wordcount

import (
	"sort"
	"strings"
)

// Frequency is a map of words to their frequency counts
type Frequency map[string]int

// WordCount counts the number of words in the given input
func WordCount(input string) Frequency {
	input = strings.ToLower(input)
	input = strip(input, []string{"!", "&", "@", "$", "%", "^", ":", "."})
	input = replace(input, []string{"\n", ",", " '", "' "}, " ")
	words := strings.Split(input, " ")
	result := make(Frequency)
	for _, word := range words {
		if word == "" {
			continue
		}
		result[word]++
	}

	return sortResult(result)
}

func strip(input string, chars []string) string {
	return replace(input, chars, "")
}

func replace(input string, chars []string, replacement string) string {
	for _, char := range chars {
		input = strings.Replace(input, char, replacement, -1)
	}
	return input
}

func sortResult(result Frequency) Frequency {
	keys := make([]string, 0, len(result))
	for key := range result {
		keys = append(keys, key)
	}
	sort.Strings(keys) //sort by key
	sorted := make(Frequency)
	for _, key := range keys {
		sorted[key] = result[key]
	}

	return sorted
}
