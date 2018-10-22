// Package proverb generates proverbs
package proverb

import "fmt"

// Proverb takes a list of words and outputs a proverb for it
func Proverb(rhyme []string) []string {
	proverb := make([]string, len(rhyme))

	for index, word := range rhyme {
		// if there is more than 2 elements left add the next line
		if index+1 < len(rhyme) {
			proverb[index] = fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[index+1])
		}
	}

	if len(rhyme) > 0 {
		// add the final line using the first word in the
		proverb[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	}

	return proverb
}
